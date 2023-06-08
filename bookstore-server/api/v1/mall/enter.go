package mall

import "bookstore/service"

type MallGroup struct {
	MallIndexApi
	MallBooksInfoApi
	MallBooksCategoryApi
	MallUserApi
	MallUserAddressApi
	MallShopCartApi
	MallOrderApi
}

var mallCarouselService = service.ServiceGroupApp.MallServiceGroup.MallCarouselService
var mallBooksInfoService = service.ServiceGroupApp.MallServiceGroup.MallBooksInfoService
var mallIndexConfigService = service.ServiceGroupApp.MallServiceGroup.MallIndexInfoService
var mallBooksCategoryService = service.ServiceGroupApp.MallServiceGroup.MallBooksCategoryService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService
var mallUserAddressService = service.ServiceGroupApp.MallServiceGroup.MallUserAddressService
var mallShopCartService = service.ServiceGroupApp.MallServiceGroup.MallShopCartService
var mallOrderService = service.ServiceGroupApp.MallServiceGroup.MallOrderService
