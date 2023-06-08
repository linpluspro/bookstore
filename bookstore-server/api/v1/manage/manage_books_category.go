package manage

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bookstore/global"
	"bookstore/model/common/enum"
	"bookstore/model/common/request"
	"bookstore/model/common/response"
	manageReq "bookstore/model/manage/request"
	manageRes "bookstore/model/manage/response"
)

type ManageBooksCategoryApi struct {
}

// CreateCategory 新建商品分类
func (g *ManageBooksCategoryApi) CreateCategory(c *gin.Context) {
	var category manageReq.MallBooksCategoryReq
	_ = c.ShouldBindJSON(&category)
	if err := mallBooksCategoryService.AddCategory(category); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateCategory 修改商品分类信息
func (g *ManageBooksCategoryApi) UpdateCategory(c *gin.Context) {
	var category manageReq.MallBooksCategoryReq
	_ = c.ShouldBindJSON(&category)
	if err := mallBooksCategoryService.UpdateCategory(category); err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败，存在相同分类", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetCategoryList 获取商品分类
func (g *ManageBooksCategoryApi) GetCategoryList(c *gin.Context) {
	var req manageReq.SearchCategoryParams
	_ = c.ShouldBindQuery(&req)
	if err, list, total := mallBooksCategoryService.SelectCategoryPage(req); err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
	} else {
		response.OkWithData(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   req.PageNumber,
			PageSize:   req.PageSize,
			TotalPage:  int(total) / req.PageSize,
		}, c)
	}
}

// GetCategory 通过id获取分类数据
func (g *ManageBooksCategoryApi) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, booksCategory := mallBooksCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
	} else {
		response.OkWithData(manageRes.BooksCategoryResponse{BooksCategory: booksCategory}, c)
	}
}

// DelCategory 设置分类失效
func (g *ManageBooksCategoryApi) DelCategory(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err, _ := mallBooksCategoryService.DeleteBooksCategoriesByIds(ids); err != nil {
		global.GVA_LOG.Error("删除失败！", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

}

// ListForSelect 用于三级分类联动效果制作
func (g *ManageBooksCategoryApi) ListForSelect(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err, booksCategory := mallBooksCategoryService.SelectCategoryById(id)
	if err != nil {
		global.GVA_LOG.Error("获取失败！", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	}
	level := booksCategory.CategoryLevel
	if level == enum.LevelThree.Code() ||
		level == enum.Default.Code() {
		response.FailWithMessage("参数异常", c)
	}
	categoryResult := make(map[string]interface{})
	if level == enum.LevelOne.Code() {
		_, levelTwoList := mallBooksCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelTwo.Code())
		if levelTwoList != nil {
			_, levelThreeList := mallBooksCategoryService.SelectByLevelAndParentIdsAndNumber(int(levelTwoList[0].CategoryId), enum.LevelThree.Code())
			categoryResult["secondLevelCategories"] = levelTwoList
			categoryResult["thirdLevelCategories"] = levelThreeList
		}
	}
	if level == enum.LevelTwo.Code() {
		_, levelThreeList := mallBooksCategoryService.SelectByLevelAndParentIdsAndNumber(id, enum.LevelThree.Code())
		categoryResult["thirdLevelCategories"] = levelThreeList
	}
	response.OkWithData(categoryResult, c)
}
