package mall

import (
	"github.com/gin-gonic/gin"

	v1 "bookstore/api/v1"
)

type MallBooksCategoryIndexRouter struct {
}

func (m *MallBooksInfoIndexRouter) InitMallBooksCategoryIndexRouter(Router *gin.RouterGroup) {
	mallBooksRouter := Router.Group("v1")
	var mallBooksCategoryApi = v1.ApiGroupApp.MallApiGroup.MallBooksCategoryApi
	{
		mallBooksRouter.GET("categories", mallBooksCategoryApi.GetBooksCategory) // 获取分类数据
	}
}
