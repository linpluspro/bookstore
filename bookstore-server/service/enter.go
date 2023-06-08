package service

import (
	"bookstore/service/example"
	"bookstore/service/mall"
	"bookstore/service/manage"
)

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
	MallServiceGroup    mall.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
