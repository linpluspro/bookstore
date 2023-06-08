package router

import (
	"bookstore/router/mall"
	"bookstore/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
