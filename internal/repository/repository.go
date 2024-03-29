package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/itrepablik/itrlog"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/webitel/sql_for_dialer/internal/model"
	"github.com/webitel/sql_for_dialer/internal/webitelClient/models"
	"reflect"
	"strings"
)

type MembersRepository interface {
	PreExecuteQuery(ctx context.Context, query string) (bool, error)
	AfterExecuteQuery(ctx context.Context, query string) (bool, error)
	GetMembers(ctx context.Context, columns []string, tableName, primaryColumn, importColumn, customFilter string) ([]map[string]interface{}, error)
	UpdateMembers(ctx context.Context, tableName, updateColumnName, primaryColumn, customFilter string) error
	CdrBulkCreate(ctx context.Context, configs *model.StatisticRequest, table string, columns []string, users []*models.EngineAttemptHistory) error
	Close() error
}

type Connect func(connString string) (MembersRepository, error)

var DriverMap = map[string]Connect{
	"psql":  NewPostgreSQLRepo,
	"mssql": NewMSSQLRepo,
	"mysql": NewMySQLRepo,
}

func NewMSSQLRepo(connString string) (MembersRepository, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return mssqlRepo{db}, nil
}

type mssqlRepo struct {
	db *sql.DB
}

func (r mssqlRepo) PreExecuteQuery(ctx context.Context, query string) (bool, error) {
	_, err := r.db.QueryContext(ctx, query)
	itrlog.Info("PreExecuteQuery: ", query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return false, err
	}
	return true, nil
}

func (r mssqlRepo) AfterExecuteQuery(ctx context.Context, query string) (bool, error) {
	_, err := r.db.QueryContext(ctx, query)
	itrlog.Info("AfterExecuteQuery: ", query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return false, err
	}
	return true, nil
}

func (r mssqlRepo) GetMembers(ctx context.Context, columns []string, tableName, primaryColumn, importColumn, customFilter string) ([]map[string]interface{}, error) {
	//TODO add custom filter
	query := fmt.Sprintf("select top 1000 %s, %s from %s where %s is null %s order by %s asc",
		primaryColumn, strings.Join(columns, ", "), tableName, importColumn, customFilter, primaryColumn)
	log.Info().Str("Select query", query)
	itrlog.Info("Select query: ", query)
	rows, err := r.db.QueryContext(ctx, query, "")
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	if rows.Err() != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	count := len(columns) + 1
	for rows.Next() {
		arr := make([]interface{}, count)
		arrPtrs := make([]interface{}, count)
		arrPtrs[0] = &arr[0]
		for i := range columns {
			arrPtrs[i+1] = &arr[i+1]
		}
		err = rows.Scan(arrPtrs...)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg(err.Error())
			return nil, err
		}
		tmp := make(map[string]interface{})
		tmp[primaryColumn] = arr[0]
		for i, col := range columns {
			tmp[col] = arr[i+1]
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (r mssqlRepo) UpdateMembers(ctx context.Context, tableName, updateColumnName, primaryColumn, customFilter string) error {
	query := fmt.Sprintf("update %s set %s = SYSDATETIME() where %s in (select top 1000 %s from %s Where %s is null %s order by %s asc) ",
		tableName, updateColumnName, primaryColumn, primaryColumn, tableName, updateColumnName, customFilter, primaryColumn)
	log.Info().Str("Update query: ", query)
	itrlog.Info("Update query: ", query)
	_, err := r.db.QueryContext(ctx, query, "")
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return err
	}
	return nil
}

func (r mssqlRepo) Close() error {
	err := r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r mssqlRepo) CdrBulkCreate(ctx context.Context, configs *model.StatisticRequest, table string, columns []string, users []*models.EngineAttemptHistory) error {
	valueArgs := []interface{}{}
	conn, err := r.db.Conn(ctx)
	err = conn.Raw(func(dbo interface{}) error {
		mssqlConn, is := dbo.(*mssql.Conn)
		if !is || mssqlConn == nil {
			return nil
		}
		bulk := mssqlConn.CreateBulk(table, columns)
		for _, user := range users {
			var isErr = true
			var errCol string
			for i, _ := range columns {
				tmp, isValid, col := getValueHelper(user, configs, i)
				valueArgs = append(valueArgs, tmp)
				if !isValid {
					isErr = isValid
					errCol = col
				}
			}
			if !isErr {
				itrlog.Warnw("Invalid column name", "Column", errCol)
				log.Warn().Str("Column", errCol).Msg("Invalid column name")
				return nil
			}
			err = bulk.AddRow(valueArgs)
			if err != nil {
				log.Err(err).Msg("Bulk failed")
				return err
			}
			valueArgs = nil
		}
		rowsCount, err := bulk.Done()
		if err != nil {
			log.Err(err).Msg("Bulk failed")
			return err
		}
		log.Info().Str("Count", fmt.Sprintf("%v", rowsCount)).Msg("Bulk success")
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func getValueHelper(user *models.EngineAttemptHistory, configs *model.StatisticRequest, i int) (interface{}, bool, string) {
	if strings.Contains(configs.Mapping[configs.Database.Table.Columns[i]], "%") {
		key := strings.Replace(configs.Mapping[configs.Database.Table.Columns[i]], "%", "", 2)
		return configs.Constants[key], true, ""
	} else if strings.Contains(configs.Mapping[configs.Database.Table.Columns[i]], ".") {
		tmp := strings.Split(configs.Mapping[configs.Database.Table.Columns[i]], ".")

		res := reflect.Indirect(reflect.ValueOf(user)).FieldByName(tmp[0]).Interface()
		for i, val := range tmp {
			if i > 0 && res != nil && !(reflect.ValueOf(res).Kind() == reflect.Ptr && reflect.ValueOf(res).IsNil()) {
				res = reflect.Indirect(reflect.ValueOf(res)).FieldByName(val).Interface()
			}
		}
		if reflect.ValueOf(res).Kind() == reflect.Ptr && reflect.ValueOf(res).IsNil() {
			res = nil
		}
		return res, true, ""
	}

	a := reflect.Indirect(reflect.ValueOf(user)).FieldByName(configs.Mapping[configs.Database.Table.Columns[i]])
	if !a.IsValid() {
		return nil, false, configs.Mapping[configs.Database.Table.Columns[i]]
	}
	return reflect.Indirect(reflect.ValueOf(user)).FieldByName(configs.Mapping[configs.Database.Table.Columns[i]]).Interface(), true, ""
}

func NewPostgreSQLRepo(connString string) (MembersRepository, error) {
	dbConn, _ := pgx.ParseConnectionString(connString)
	db, err := pgx.Connect(dbConn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}
	return psqlRepo{db}, nil
}

type psqlRepo struct {
	db *pgx.Conn
}

func (r psqlRepo) PreExecuteQuery(ctx context.Context, query string) (bool, error) {
	err := r.db.QueryRow(query).Scan()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r psqlRepo) AfterExecuteQuery(ctx context.Context, query string) (bool, error) {
	err := r.db.QueryRow(query).Scan()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r psqlRepo) GetMembers(ctx context.Context, columns []string, tableName, primaryColumn, importColumn, customFilter string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("select %s, %s from %s where %s IS NULL;", primaryColumn, strings.Join(columns, ", "), tableName, importColumn)

	r.db.QueryRow(query).Scan()
	return nil, nil
}

func (r psqlRepo) UpdateMembers(ctx context.Context, tableName, updateColumnName, primaryColumn, customFilter string) error {
	return nil
}

func (r psqlRepo) Close() error {
	return nil
}

func (r psqlRepo) CdrBulkCreate(ctx context.Context, configs *model.StatisticRequest, table string, columns []string, users []*models.EngineAttemptHistory) error {
	return nil
}

func NewMySQLRepo(connString string) (MembersRepository, error) {
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return mysqlRepo{db}, nil
}

type mysqlRepo struct {
	db *sql.DB
}

func (r mysqlRepo) PreExecuteQuery(ctx context.Context, query string) (bool, error) {
	_, err := r.db.ExecContext(ctx, query)
	itrlog.Info("PreExecuteQuery: ", query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return false, err
	}
	return true, nil
}

func (r mysqlRepo) AfterExecuteQuery(ctx context.Context, query string) (bool, error) {
	_, err := r.db.ExecContext(ctx, query)
	itrlog.Info("AfterExecuteQuery: ", query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return false, err
	}
	return true, nil
}

func (r mysqlRepo) GetMembers(ctx context.Context, columns []string, tableName, primaryColumn, importColumn, customFilter string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("select %s, %s from %s where %s is null %s order by %s asc LIMIT 1000",
		primaryColumn, strings.Join(columns, ", "), tableName, importColumn, customFilter, primaryColumn)
	log.Info().Str("Select query", query)
	itrlog.Info("Select query: ", query)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	if rows.Err() != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return nil, err
	}

	count := len(columns) + 1
	for rows.Next() {
		arr := make([]interface{}, count)
		arrPtrs := make([]interface{}, count)
		arrPtrs[0] = &arr[0]
		for i := range columns {
			arrPtrs[i+1] = &arr[i+1]
		}
		err = rows.Scan(arrPtrs...)
		if err != nil {
			itrlog.Error(err.Error())
			log.Err(err).Msg(err.Error())
			return nil, err
		}
		tmp := make(map[string]interface{})
		tmp[primaryColumn] = arr[0]
		for i, col := range columns {
			tmp[col] = arr[i+1]
		}
		result = append(result, tmp)
	}
	return result, nil
}

func (r mysqlRepo) UpdateMembers(ctx context.Context, tableName, updateColumnName, primaryColumn, customFilter string) error {
	query := fmt.Sprintf("update %s set %s = current_timestamp where %s in (SELECT * FROM (select %s from %s Where %s is null %s order by %s asc LIMIT 1000) as sq) ",
		tableName, updateColumnName, primaryColumn, primaryColumn, tableName, updateColumnName, customFilter, primaryColumn)
	log.Info().Str("Update query: ", query)
	itrlog.Info("Update query: ", query)
	_, err := r.db.ExecContext(ctx, query)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return err
	}
	return nil
}

func (r mysqlRepo) Close() error {
	err := r.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (r mysqlRepo) CdrBulkCreate(ctx context.Context, configs *model.StatisticRequest, table string, columns []string, users []*models.EngineAttemptHistory) error {

	conn, err := r.db.Conn(ctx)
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
	}
	for _, user := range users {
		//valueStrings := "(?"+strings.Repeat(",?", len(columns)-1)+")"
		valueArgs := []string{}

		for i, _ := range columns {
			tmp, isValid, col := getValueHelper(user, configs, i)
			if !isValid {
				itrlog.Warnw("Invalid column name", "Column", col)
				log.Warn().Str("Column", col).Msg("Invalid column name")
				return err
			}
			if tmp != nil {
				valueArgs = append(valueArgs, fmt.Sprintf("\"%v\"", tmp))
			} else {
				valueArgs = append(valueArgs, "\"\"")
			}
		}
		smt := `INSERT INTO %s(%s) VALUES (%s)`
		smt = fmt.Sprintf(smt, table, strings.Join(columns, ","), strings.Join(valueArgs, ","))

		_, err = tx.ExecContext(ctx, smt)

		if err != nil {
			tx.Rollback()
			itrlog.Error(err.Error())
			log.Err(err).Msg(err.Error())
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		itrlog.Error(err.Error())
		log.Err(err).Msg(err.Error())
		return err
	}
	return nil
}
