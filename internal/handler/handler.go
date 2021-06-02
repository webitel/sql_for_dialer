package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/rs/zerolog/log"
	"github.com/webitel/sql_for_dialer/internal/model"
	"github.com/webitel/sql_for_dialer/internal/repository"
	apiclient "github.com/webitel/sql_for_dialer/internal/webitelClient/client"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/client/member_service"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetMembers(w http.ResponseWriter, req *http.Request) {
	var memberReq []*model.MemberRequest
	err := json.NewDecoder(req.Body).Decode(&memberReq)
	if err != nil {
		log.Err(err).Msg(err.Error())
		return
	}
	defer req.Body.Close()
	configs := memberReq[0]
	repo, err := repository.DriverMap[configs.Database.DbType](configs.Database.ConnectionString)
	if err != nil {
		log.Err(err).Msg("")
		return
	}
	defer repo.Close()
	if configs.Database.BeforeExecute != "" {
		_, err := repo.PreExecuteQuery(req.Context(), configs.Database.BeforeExecute)
		if err != nil {
			log.Err(err).Msg("")
			return
		}
	}
	for {
		members, err := repo.GetMembers(req.Context(), configs.Database.Table.Columns, configs.Database.Table.Name, configs.Database.Table.PrimaryCol, configs.Database.Table.ImportDateCol, configs.Database.CustomSqlFilter)
		if err != nil {
			log.Err(err).Msg("")
			return
		}

		if len(members) == 0 {
			break
		}
		log.Info().Str("count", fmt.Sprintf("%v", len(members))).Msg("Rows selected")

		r := httptransport.New(configs.Webitel.Host, configs.Webitel.BasePath, apiclient.DefaultSchemes)
		r.DefaultAuthentication = httptransport.APIKeyAuth("X-Webitel-Access", "header", configs.Webitel.Token)
		//r.SetDebug(true)
		memberService := member_service.New(r, strfmt.Default)

		webitelItems := make([]*models.EngineCreateMemberBulkItem, 0, len(members))
		for _, item := range members {
			communications := make([]*models.EngineMemberCommunicationCreateRequest, 0, len(configs.Mapping.Destinations))
			for index, value := range configs.Mapping.Destinations {
				phoneType := &models.EngineLookup{}
				if strings.Contains(configs.Mapping.PhoneTypes[index], "%") {
					key := strings.Replace(configs.Mapping.PhoneTypes[index], "%", "", 2)
					phoneType = &models.EngineLookup{
						ID: configs.Constants[key],
					}
				} else {
					phoneType = &models.EngineLookup{
						ID: item[configs.Mapping.PhoneTypes[index]].(string),
					}
				}

				communications = append(communications, &models.EngineMemberCommunicationCreateRequest{
					Destination: item[value].(string),
					Type:        phoneType,
				})
			}

			name, _ := item[configs.Mapping.Name].(string)

			webitelItems = append(webitelItems, &models.EngineCreateMemberBulkItem{
				Communications: communications,
				Name:           name,
				Variables:      configs.Mapping.Variables,
			})
		}

		webitelReq := &models.EngineCreateMemberBulkRequest{
			Items:   webitelItems,
			QueueID: strconv.Itoa(configs.Webitel.QueueId),
		}

		result, err := memberService.CreateMemberBulk(&member_service.CreateMemberBulkParams{
			Body:    webitelReq,
			QueueID: strconv.Itoa(configs.Webitel.QueueId),
			Context: context.Background(),
		}, r.DefaultAuthentication)

		if err != nil {
			log.Err(err).Msg("")
			return
		}
		if len(result.GetPayload().Ids) != 0 {
			log.Info().Str("count", fmt.Sprintf("%v", len(result.GetPayload().Ids))).Msg("Members inserted")
			err := repo.UpdateMembers(req.Context(), configs.Database.Table.Name, configs.Database.Table.ImportDateCol, configs.Database.Table.PrimaryCol)
			if err != nil {
				log.Err(err).Msg("")
				return
			}
		}
	}

	if configs.Database.AfterExecute != "" {
		_, err := repo.AfterExecuteQuery(req.Context(), configs.Database.AfterExecute)
		if err != nil {
			return
		}
	}

}

func SetMemberStat(w http.ResponseWriter, req *http.Request) {
	buf, _ := ioutil.ReadAll(req.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	var statsticReq []*model.StatisticRequest
	err := json.NewDecoder(rdr1).Decode(&statsticReq)
	if err != nil {
		log.Err(err).Msg("Bad request")
		w.WriteHeader(400)
		return
	}
	log.Info().Msg(fmt.Sprintf("%q", rdr2))
	defer req.Body.Close()
	configs := statsticReq[0]
	r := httptransport.New(configs.Webitel.Host, configs.Webitel.BasePath, apiclient.DefaultSchemes)
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
			return
		}

		err = repo.CdrBulkCreate(context.Background(), configs, configs.Database.Table.Name, configs.Database.Table.Columns, result.Payload.Items)
		if err != nil {
			return
		}
		if !result.Payload.Next {
			break
		}
		i++
	}
	return
}
