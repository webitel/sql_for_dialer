package model

/* Model for Webitel members request */

type MemberRequest struct {
	Webitel   Webitel           `json:"webitel,omitempty"`
	Database  Database          `json:"database,omitempty"`
	Mapping   Mapping           `json:"mapping,omitempty"`
	Constants map[string]string `json:"constants,omitempty"`
}

type Webitel struct {
	Host     string `json:"host,omitempty"`
	BasePath string `json:"basePath,omitempty"`
	Token    string `json:"token,omitempty"`
	QueueId  int    `json:"queueId,omitempty"`

	MembersTTL string `json:"membersTTL,omitempty"`
}

type Database struct {
	ConnectionString string `json:"connectionString,omitempty"`
	DbType           string `json:"dbType,omitempty"`
	Table            Table  `json:"table,omitempty"`
	CustomSqlFilter  string `json:"customSqlFilter,omitempty"`
	BeforeExecute    string `json:"beforeExecute,omitempty"`
	AfterExecute     string `json:"afterExecute,omitempty"`
}

type Table struct {
	Name          string   `json:"name,omitempty"`
	Columns       []string `json:"columns,omitempty"`
	PrimaryCol    string   `json:"primaryCol,omitempty"`
	ImportDateCol string   `json:"importDateCol,omitempty"`
}

type Mapping struct {
	Name         string            `json:"name,omitempty"`
	Destinations []string          `json:"destinations,omitempty"`
	PhoneTypes   []string          `json:"phoneTypes,omitempty"`
	Variables    map[string]string `json:"variables,omitempty"`
}

/* Model for Webitel statistic request */

type StatisticRequest struct {
	Webitel   StatisticWebitel  `json:"webitel,omitempty"`
	Database  StatisticDatabase `json:"database,omitempty"`
	Mapping   map[string]string `json:"mapping,omitempty"`
	Constants map[string]string `json:"constants,omitempty"`
}

type StatisticWebitel struct {
	Host     string `json:"host,omitempty"`
	BasePath string `json:"basePath,omitempty"`
	Token    string `json:"token,omitempty"`
	QueueId  int    `json:"queueId,omitempty"`

	MembersTTL string `json:"membersTTL,omitempty"`
}

type StatisticDatabase struct {
	ConnectionString string         `json:"connectionString,omitempty"`
	DbType           string         `json:"dbType,omitempty"`
	PackSize         int32          `json:"packSize,omitempty"`
	Table            StatisticTable `json:"table,omitempty"`
}

type StatisticTable struct {
	Name    string   `json:"name,omitempty"`
	Columns []string `json:"columns,omitempty"`
}
