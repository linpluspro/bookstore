package request

import (
	"bookstore/model/common/request"
	"bookstore/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}
