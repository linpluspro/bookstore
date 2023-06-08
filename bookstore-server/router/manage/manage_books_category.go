package manage

import (
	"github.com/gin-gonic/gin"

	v1 "bookstore/api/v1"
	"bookstore/middleware"
)

type ManageBooksCategoryRouter struct {
}

func (r *ManageBooksCategoryRouter) InitManageBooksCategoryRouter(Router *gin.RouterGroup) {
	booksCategoryRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var booksCategoryApi = v1.ApiGroupApp.ManageApiGroup.ManageBooksCategoryApi
	{
		booksCategoryRouter.POST("categories", booksCategoryApi.CreateCategory)
		booksCategoryRouter.PUT("categories", booksCategoryApi.UpdateCategory)
		booksCategoryRouter.GET("categories", booksCategoryApi.GetCategoryList)
		booksCategoryRouter.GET("categories/:id", booksCategoryApi.GetCategory)
		booksCategoryRouter.DELETE("categories", booksCategoryApi.DelCategory)
		booksCategoryRouter.GET("categories4Select", booksCategoryApi.ListForSelect)
	}
}
