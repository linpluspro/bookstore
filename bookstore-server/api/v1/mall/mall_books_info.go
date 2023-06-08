package mall

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bookstore/global"
	"bookstore/model/common/response"
	mallReq "bookstore/model/mall/request"
)

type MallBooksInfoApi struct {
}

// BooksSearch 商品搜索
func (m *MallBooksInfoApi) BooksSearch(c *gin.Context) {
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	booksCategoryId, _ := strconv.Atoi(c.Query("booksCategoryId"))
	keyword := c.Query("keyword")
	orderBy := c.Query("orderBy")
	if err, list, total := mallBooksInfoService.MallBooksListBySearch(pageNumber, booksCategoryId, keyword, orderBy); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   10,
		}, "获取成功", c)
	}
}

func (m *MallBooksInfoApi) BooksDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, booksInfo := mallBooksInfoService.GetMallBooksInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	} else {
		response.OkWithData(booksInfo, c)
	}

}

func (m *MallBooksInfoApi) Comment(c *gin.Context) {
	var req mallReq.CreateBookCommentParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}

	token := c.GetHeader("token")
	if err := mallUserService.CreateBookComment(token, req); err != nil {
		global.GVA_LOG.Error("发布评论失败", zap.Error(err))
		response.FailWithMessage("发布评论失败"+err.Error(), c)
	} else {
		response.OkWithMessage("发布评论成功", c)
	}
}

func (m *MallBooksInfoApi) CommentReply(c *gin.Context) {
	var req mallReq.CreateBookCommentParam
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage("参数错误"+err.Error(), c)
		return
	}

	token := c.GetHeader("token")
	if err := mallUserService.CreateBookComment(token, req); err != nil {
		global.GVA_LOG.Error("回复评论失败", zap.Error(err))
		response.FailWithMessage("回复评论失败"+err.Error(), c)
	} else {
		response.OkWithMessage("回复评论成功", c)
	}
}
