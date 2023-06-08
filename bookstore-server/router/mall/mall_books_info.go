package mall

import (
	"github.com/gin-gonic/gin"

	v1 "bookstore/api/v1"
)

type MallBooksInfoIndexRouter struct {
}

func (m *MallBooksInfoIndexRouter) InitMallBooksInfoIndexRouter(Router *gin.RouterGroup) {
	mallBooksRouter := Router.Group("v1")
	var mallBooksInfoApi = v1.ApiGroupApp.MallApiGroup.MallBooksInfoApi
	{
		mallBooksRouter.GET("/search", mallBooksInfoApi.BooksSearch)               // 商品搜索
		mallBooksRouter.GET("/books/detail/:id", mallBooksInfoApi.BooksDetail)     // 商品详情
		mallBooksRouter.POST("/books/comment", mallBooksInfoApi.Comment)           // 评论
		mallBooksRouter.POST("/books/commentReply", mallBooksInfoApi.CommentReply) // 评论回复
	}
}
