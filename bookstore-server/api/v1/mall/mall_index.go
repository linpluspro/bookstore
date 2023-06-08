package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"bookstore/global"
	"bookstore/model/common/enum"
	"bookstore/model/common/response"
)

type MallIndexApi struct {
}

// MallIndexInfo 加载首页信息
func (m *MallIndexApi) MallIndexInfo(c *gin.Context) {
	err, _, mallCarouseInfo := mallCarouselService.GetCarouselsForIndex(5)
	if err != nil {
		global.GVA_LOG.Error("轮播图获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("轮播图获取失败", c)
	}
	err, hotBooks := mallIndexConfigService.GetConfigBooksForIndex(enum.IndexBooksHot.Code(), 4)
	if err != nil {
		global.GVA_LOG.Error("热销图书获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("热销图书获取失败", c)
	}
	err, newBooks := mallIndexConfigService.GetConfigBooksForIndex(enum.IndexBooksNew.Code(), 5)
	if err != nil {
		global.GVA_LOG.Error("新书获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("新书获取失败", c)
	}
	err, recommendBooks := mallIndexConfigService.GetConfigBooksForIndex(enum.IndexBooksRecommond.Code(), 10)
	if err != nil {
		global.GVA_LOG.Error("推荐图书获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("推荐图书获取失败", c)
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = mallCarouseInfo
	indexResult["hotBooks"] = hotBooks
	indexResult["newBooks"] = newBooks
	indexResult["recommendBooks"] = recommendBooks
	response.OkWithData(indexResult, c)
}
