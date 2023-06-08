package manage

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bookstore/global"
	"bookstore/model/common/request"
	"bookstore/model/common/response"
	manageReq "bookstore/model/manage/request"
)

type ManageCarouselApi struct {
}

func (m *ManageCarouselApi) CreateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselAddParam
	_ = c.ShouldBindJSON(&req)
	if err := mallCarouselService.CreateCarousel(req); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (m *ManageCarouselApi) DeleteCarousel(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := mallCarouselService.DeleteCarousel(ids); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (m *ManageCarouselApi) UpdateCarousel(c *gin.Context) {
	var req manageReq.MallCarouselUpdateParam
	_ = c.ShouldBindJSON(&req)
	if err := mallCarouselService.UpdateCarousel(req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCarousel 用id查询MallCarousel
func (m *ManageCarouselApi) FindCarousel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, mallCarousel := mallCarouselService.GetCarousel(id); err != nil {
		global.GVA_LOG.Error("查询失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(mallCarousel, c)
	}
}

// GetCarouselList 分页获取MallCarousel列表
func (m *ManageCarouselApi) GetCarouselList(c *gin.Context) {
	var pageInfo manageReq.MallCarouselSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallCarouselService.GetCarouselInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!"+err.Error(), zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
