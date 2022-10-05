package handler

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/itrepablik/itrlog"
	"github.com/rs/zerolog/log"
	"github.com/webitel/sql_for_dialer/internal/model"
	"github.com/webitel/sql_for_dialer/internal/repository"
	apiclient "github.com/webitel/sql_for_dialer/internal/webitelClient/client"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/client/bucket_service"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/client/communication_type_service"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/client/member_service"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
	"io"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

var skipTlsClient = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	},
}

func GetLogs(w http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	dat, err := ioutil.ReadFile(fmt.Sprintf("logs/get_members_log_%s.log", currentTime.Format("2006-01-02")))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "{\n\"error\" : \"file not found\"\n}\n")
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(dat))
}

func GetCurrentVersion(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "{\n\"version\" : \"1.0.9\"\n}\n")
}

func GetMembers(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			itrlog.Error(fmt.Sprintf("stacktrace from panic: \n" + string(debug.Stack())))
		}
	}()
	itrlog.SetLogInit(50, 90, "logs", "get_members_log_")
	var memberReq []*model.MemberRequest
	err := json.NewDecoder(req.Body).Decode(&memberReq)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		io.WriteString(w, fmt.Sprintf("{\n\"error\" : %q\n}", err.Error()))
		return
	}
	defer req.Body.Close()
	configs := memberReq[0]
	repo, err := repository.DriverMap[configs.Database.DbType](configs.Database.ConnectionString)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg("Err connection to DB or incorrect connection string")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		io.WriteString(w, "{\n\"error\" : \"Unable connect to DB or incorrect connection string\"\n}\n")
		return
	}
	itrlog.Info("Success connection to DB")
	defer repo.Close()
	if configs.Database.BeforeExecute != "" {
		_, err := repo.PreExecuteQuery(req.Context(), configs.Database.BeforeExecute)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}
	}
	r := httptransport.NewWithClient(configs.Webitel.Host, configs.Webitel.BasePath, apiclient.DefaultSchemes, skipTlsClient)
	r.DefaultAuthentication = httptransport.APIKeyAuth("X-Webitel-Access", "header", configs.Webitel.Token)

	deleteMemberService := member_service.New(r, strfmt.Default)
	if configs.Webitel.DeleteMembers {
		deleteRes, err := deleteMemberService.DeleteMembers(&member_service.DeleteMembersParams{
			Body:    nil,
			QueueID: fmt.Sprintf("%v", configs.Webitel.QueueId),
			Context: context.Background(),
		}, r.DefaultAuthentication)

		if err != nil {
			itrlog.Error(fmt.Sprintf("Error on delete members in QueueId: %v. Error: %v", configs.Webitel.QueueId, err.Error()))
			log.Err(err).Msg(fmt.Sprintf("Error on delete members in QueueId: %v", configs.Webitel.QueueId))
			return
		}
		itrlog.Info(fmt.Sprintf("Success delete %v members in QueueId: %v", len(deleteRes.GetPayload().Items), configs.Webitel.QueueId))
		log.Info().Msg(fmt.Sprintf("Success delete %v members in QueueId: %v", len(deleteRes.GetPayload().Items), configs.Webitel.QueueId))
	}

	for {
		members, err := repo.GetMembers(req.Context(), configs.Database.Table.Columns, configs.Database.Table.Name, configs.Database.Table.PrimaryCol, configs.Database.Table.ImportDateCol, configs.Database.CustomSqlFilter)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}

		if len(members) == 0 {
			itrlog.Info(fmt.Sprintf("Rows selected. Count: %v", len(members)))
			log.Info().Msg("Count 0")
			break
		}
		itrlog.Info(fmt.Sprintf("Rows selected. Count: %v", len(members)))
		log.Info().Str("count", fmt.Sprintf("%v", len(members))).Msg("Rows selected")

		var size int32 = 20

		communicationTypeService := communication_type_service.New(r, strfmt.Default)

		commTypes, err := communicationTypeService.SearchCommunicationType(&communication_type_service.SearchCommunicationTypeParams{
			Page:    nil,
			Size:    &size,
			Context: context.Background(),
		}, r.DefaultAuthentication)

		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}
		if len(commTypes.GetPayload().Items) == 0 {
			log.Warn().Msg("There is no communication types")
			itrlog.Warn("There is no communication types")
		}

		itrlog.Info("Success connection to Webitel")

		webitelItems := make([]*models.EngineCreateMemberBulkItem, 0, len(members))
		for _, item := range members {
			communications := make([]*models.EngineMemberCommunicationCreateRequest, 0, len(configs.Mapping.Destinations))
			for index, value := range configs.Mapping.Destinations {
				dest := ""
				switch v := item[value].(type) {
				case int, int64, int32:
					dest = fmt.Sprintf("%d", v)
				case string:
					dest = fmt.Sprintf("%s", v)
				case []uint8:
					dest = fmt.Sprintf("%s", string(v))
				case time.Time:
					dest = fmt.Sprintf("%d", v.UnixNano()/int64(time.Millisecond))
				}

				if dest != "" {
					phoneType := &models.EngineLookup{}
					if strings.Contains(configs.Mapping.PhoneTypes[index], "%") {
						key := strings.Replace(configs.Mapping.PhoneTypes[index], "%", "", 2)
						phoneType = &models.EngineLookup{
							ID: configs.Constants[key],
						}
					} else {
						ptCode := ""
						switch v := item[configs.Mapping.PhoneTypes[index]].(type) {
						case int, int64, int32:
							ptCode = fmt.Sprintf("%d", v)
						case string:
							ptCode = fmt.Sprintf("%s", v)
						case []uint8:
							ptCode = fmt.Sprintf("%s", string(v))
						case time.Time:
							ptCode = fmt.Sprintf("%d", v.UnixNano()/int64(time.Millisecond))
						}
						ptId := ""
						for _, item := range commTypes.Payload.Items {
							if item.Code == ptCode {
								ptId = item.ID
							}
						}
						phoneType = &models.EngineLookup{
							ID: ptId,
						}
					}
					communications = append(communications, &models.EngineMemberCommunicationCreateRequest{
						Destination: dest,
						Type:        phoneType,
					})
				}
			}
			vers := make(map[string]string)
			for i, val := range configs.Mapping.Variables {
				if strings.Contains(val, "%") {
					key := strings.Replace(val, "%", "", 2)
					vers[i] = key
				} else {
					switch v := item[val].(type) {
					case int, int64, int32:
						vers[i] = fmt.Sprintf("%d", v)
					case string:
						vers[i] = fmt.Sprintf("%s", v)
					case []uint8:
						vers[i] = fmt.Sprintf("%s", string(v))
					case time.Time:
						vers[i] = fmt.Sprintf("%d", v.UnixNano()/int64(time.Millisecond))
					default:
						vers[i] = ""
					}
				}
			}

			name := ""
			switch v := item[configs.Mapping.Name].(type) {
			case int, int64, int32:
				name = fmt.Sprintf("%d", v)
			case string:
				name = fmt.Sprintf("%s", v)
			case []uint8:
				name = fmt.Sprintf("%s", string(v))
			case time.Time:
				name = fmt.Sprintf("%d", v.UnixNano()/int64(time.Millisecond))
			default:
				name = ""
			}

			bucket := &models.EngineLookup{}
			bucketName := ""
			switch v := item[configs.Mapping.Bucket].(type) {
			case int, int64, int32:
				bucketName = fmt.Sprintf("%d", v)
			case string:
				bucketName = fmt.Sprintf("%s", v)
			case []uint8:
				bucketName = fmt.Sprintf("%s", string(v))
			case time.Time:
				bucketName = fmt.Sprintf("%d", v.UnixNano()/int64(time.Millisecond))
			default:
				bucketName = ""
			}

			var sizeb int32 = 1

			bucketService := bucket_service.New(r, strfmt.Default)

			buckets, _ := bucketService.SearchBucket(&bucket_service.SearchBucketParams{
				Q:       &bucketName,
				Page:    nil,
				Size:    &sizeb,
				Context: context.Background(),
			}, r.DefaultAuthentication)
			buckId := ""
			for _, item := range buckets.Payload.Items {
				if item.Name == bucketName {
					buckId = item.ID
				}
			}
			bucket = &models.EngineLookup{
				ID: buckId,
			}

			timeToLoad, _ := time.ParseDuration(configs.Webitel.MembersTTL)
			expireAt := fmt.Sprintf("%d", time.Now().Add(timeToLoad).UnixNano()/int64(time.Millisecond))
			webitelItems = append(webitelItems, &models.EngineCreateMemberBulkItem{
				ExpireAt:       expireAt,
				Communications: communications,
				Name:           name,
				Variables:      vers,
				Bucket:         bucket,
			})
		}

		itrlog.Info("Success parse DB rows")

		webitelReq := &models.EngineCreateMemberBulkRequest{
			Items:   webitelItems,
			QueueID: fmt.Sprintf("%v", configs.Webitel.QueueId),
		}

		memberService := member_service.New(r, strfmt.Default)

		result, err := memberService.CreateMemberBulk(&member_service.CreateMemberBulkParams{
			Body:    webitelReq,
			QueueID: fmt.Sprintf("%v", configs.Webitel.QueueId),
			Context: context.Background(),
		}, r.DefaultAuthentication)

		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}
		if len(result.GetPayload().Ids) != 0 {
			itrlog.Info(fmt.Sprintf("Members inserted. Count: %v", len(result.GetPayload().Ids)))
			log.Info().Str("count", fmt.Sprintf("%v", len(result.GetPayload().Ids))).Msg("Members inserted")
			err := repo.UpdateMembers(req.Context(), configs.Database.Table.Name, configs.Database.Table.ImportDateCol, configs.Database.Table.PrimaryCol)
			if err != nil {
				itrlog.Error(err.Error())
				log.Err(err).Msg("")
				return
			}
		}
	}

	if configs.Database.AfterExecute != "" {
		_, err := repo.AfterExecuteQuery(req.Context(), configs.Database.AfterExecute)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}
	}
}

func SetMemberStat(w http.ResponseWriter, req *http.Request) {
	itrlog.SetLogInit(50, 90, "logs", "set_members_stat_log_")
	buf, _ := ioutil.ReadAll(req.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	var statsticReq []*model.StatisticRequest
	err := json.NewDecoder(rdr1).Decode(&statsticReq)
	if err != nil {
		itrlog.Errorw("Bad request", "error", err.Error())
		log.Err(err).Msg("Bad request")
		w.WriteHeader(400)
		return
	}
	itrlog.Infow("Request body", "log_time", time.Now().Format(itrlog.LogTimeFormat), "body", fmt.Sprintf("%q", rdr2))
	log.Info().Msg(fmt.Sprintf("%q", rdr2))
	defer req.Body.Close()
	configs := statsticReq[0]
	r := httptransport.NewWithClient(configs.Webitel.Host, configs.Webitel.BasePath, apiclient.DefaultSchemes, skipTlsClient)
	r.DefaultAuthentication = httptransport.APIKeyAuth("X-Webitel-Access", "header", configs.Webitel.Token)

	queueId := []string{strconv.Itoa(configs.Webitel.QueueId)}
	timeFrom := "1000000000000"
	timeTo := strconv.FormatInt(time.Now().Unix()*1000, 10)
	var size int32 = 1000
	if configs.Database.PackSize != 0 {
		size = configs.Database.PackSize
	}
	memberService := member_service.New(r, strfmt.Default)
	var i int32 = 1
	repo, err := repository.DriverMap[configs.Database.DbType](configs.Database.ConnectionString)
	if err != nil {
		itrlog.Errorw("Connect to DB", "error", err.Error())
		log.Err(err).Msg("Connect to DB")
		return
	}
	defer repo.Close()
	for {
		result, err := memberService.SearchAttemptsHistory(&member_service.SearchAttemptsHistoryParams{
			Page:         &i,
			JoinedAtFrom: &timeFrom,
			JoinedAtTo:   &timeTo,
			Size:         &size,
			QueueID:      queueId,
			Context:      context.Background(),
		}, r.DefaultAuthentication)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}

		err = repo.CdrBulkCreate(context.Background(), configs, configs.Database.Table.Name, configs.Database.Table.Columns, result.Payload.Items)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg("")
			return
		}
		if !result.Payload.Next {
			break
		}
		i++
	}
	return
}
