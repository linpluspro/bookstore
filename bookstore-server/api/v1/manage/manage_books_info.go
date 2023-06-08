package manage

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bookstore/global"
	"bookstore/model/common/request"
	"bookstore/model/common/response"
	"bookstore/model/manage"
	manageReq "bookstore/model/manage/request"
)

type ManageBooksInfoApi struct {
}

func (m *ManageBooksInfoApi) CreateBooksInfo(c *gin.Context) {
	var mallBooksInfo manageReq.BooksInfoAddParam
	_ = c.ShouldBindJSON(&mallBooksInfo)
	if err := mallBooksInfoService.CreateMallBooksInfo(mallBooksInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败!"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBooksInfo 删除MallBooksInfo
func (m *ManageBooksInfoApi) DeleteBooksInfo(c *gin.Context) {
	var mallBooksInfo manage.MallBooksInfo
	_ = c.ShouldBindJSON(&mallBooksInfo)
	if err := mallBooksInfoService.DeleteMallBooksInfo(mallBooksInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// ChangeBooksInfoByIds 批量删除MallBooksInfo
func (m *ManageBooksInfoApi) ChangeBooksInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	sellStatus := c.Param("status")
	if err := mallBooksInfoService.ChangeMallBooksInfoByIds(IDS, sellStatus); err != nil {
		global.GVA_LOG.Error("修改商品状态失败!", zap.Error(err))
		response.FailWithMessage("修改商品状态失败"+err.Error(), c)
	} else {
		response.OkWithMessage("修改商品状态成功", c)
	}
}

// UpdateBooksInfo 更新MallBooksInfo
func (m *ManageBooksInfoApi) UpdateBooksInfo(c *gin.Context) {
	var mallBooksInfo manageReq.BooksInfoUpdateParam
	_ = c.ShouldBindJSON(&mallBooksInfo)
	if err := mallBooksInfoService.UpdateMallBooksInfo(mallBooksInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBooksInfo 用id查询MallBooksInfo
func (m *ManageBooksInfoApi) FindBooksInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, booksInfo := mallBooksInfoService.GetMallBooksInfo(id)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}
	booksInfoRes := make(map[string]interface{})
	booksInfoRes["books"] = booksInfo
	if _, thirdCategory := mallBooksCategoryService.SelectCategoryById(*booksInfo.BooksCategoryId); thirdCategory != (manage.MallBooksCategory{}) {
		booksInfoRes["thirdCategory"] = thirdCategory
		if _, secondCategory := mallBooksCategoryService.SelectCategoryById(thirdCategory.ParentId); secondCategory != (manage.MallBooksCategory{}) {
			booksInfoRes["secondCategory"] = secondCategory
			if _, firstCategory := mallBooksCategoryService.SelectCategoryById(secondCategory.ParentId); firstCategory != (manage.MallBooksCategory{}) {
				booksInfoRes["firstCategory"] = firstCategory
			}
		}
	}
	response.OkWithData(booksInfoRes, c)

}

// GetBooksInfoList 分页获取MallBooksInfo列表
func (m *ManageBooksInfoApi) GetBooksInfoList(c *gin.Context) {
	var pageInfo manageReq.MallBooksInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	booksName := c.Query("booksName")
	booksSellStatus := c.GetInt("booksSellStatus")
	if err, list, total := mallBooksInfoService.GetMallBooksInfoInfoList(pageInfo, booksName, booksSellStatus); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
