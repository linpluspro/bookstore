package request

import (
	"bookstore/model/common/request"
	"bookstore/model/manage"
)

type MallIndexConfigSearch struct {
	manage.MallIndexConfig
	request.PageInfo
}

type MallIndexConfigAddParams struct {
	ConfigName  string `json:"configName"`
	ConfigType  int    `json:"configType"`
	BooksId     int    `json:"booksId"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigRank  int    `json:"configRank"`
}

type MallIndexConfigUpdateParams struct {
	ConfigId    int    `json:"configId"`
	ConfigName  string `json:"configName"`
	RedirectUrl string `json:"redirectUrl"`
	ConfigType  int    `json:"configType"`
	BooksId     int    `json:"booksId"`
	ConfigRank  int    `json:"configRank"`
}
