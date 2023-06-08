package manage

import (
	"github.com/gin-gonic/gin"

	v1 "bookstore/api/v1"
)

type ManageBooksInfoRouter struct {
}

func (m *ManageBooksInfoRouter) InitManageBooksInfoRouter(Router *gin.RouterGroup) {
	mallBooksInfoRouter := Router.Group("v1")
	var mallBooksInfoApi = v1.ApiGroupApp.ManageApiGroup.ManageBooksInfoApi
	{
		mallBooksInfoRouter.POST("books", mallBooksInfoApi.CreateBooksInfo)                    // 新建MallBooksInfo
		mallBooksInfoRouter.DELETE("deleteMallBooksInfo", mallBooksInfoApi.DeleteBooksInfo)    // 删除MallBooksInfo
		mallBooksInfoRouter.PUT("books/status/:status", mallBooksInfoApi.ChangeBooksInfoByIds) // 上下架
		mallBooksInfoRouter.PUT("books", mallBooksInfoApi.UpdateBooksInfo)                     // 更新MallBooksInfo
		mallBooksInfoRouter.GET("books/:id", mallBooksInfoApi.FindBooksInfo)                   // 根据ID获取MallBooksInfo
		mallBooksInfoRouter.GET("books/list", mallBooksInfoApi.GetBooksInfoList)               // 获取MallBooksInfo列表
	}
}
