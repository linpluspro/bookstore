package v1

import (
	"bookstore/api/v1/mall"
	"bookstore/api/v1/manage"
)

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
	MallApiGroup   mall.MallGroup
}

var ApiGroupApp = new(ApiGroup)
