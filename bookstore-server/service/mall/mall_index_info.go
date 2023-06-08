package mall

import (
	"bookstore/global"
	"bookstore/model/mall/response"
	"bookstore/model/manage"
)

type MallIndexInfoService struct {
}

// GetConfigBooksForIndex 首页返回相关IndexConfig
func (m *MallIndexInfoService) GetConfigBooksForIndex(configType int, num int) (err error, list interface{}) {
	var indexConfigs []manage.MallIndexConfig
	err = global.GVA_DB.Where("config_type = ?", configType).Where("is_deleted = 0").Order("config_rank desc").Limit(num).Find(&indexConfigs).Error
	if err != nil {
		return
	}
	// 获取商品id
	var ids []int
	for _, indexConfig := range indexConfigs {
		ids = append(ids, indexConfig.BooksId)
	}
	// 获取商品信息
	var booksList []manage.MallBooksInfo
	err = global.GVA_DB.Where("books_id in ?", ids).Find(&booksList).Error
	var indexBooksList []response.MallIndexConfigBooksResponse
	// 超出30个字符显示....
	for _, indexBooks := range booksList {
		res := response.MallIndexConfigBooksResponse{
			BooksId:       *indexBooks.BooksId,
			BooksName:     indexBooks.BooksName,
			BooksIntro:    indexBooks.BooksIntro,
			BooksCoverImg: indexBooks.BooksCoverImg,
			SellingPrice:  *indexBooks.SellingPrice,
			Tag:           indexBooks.Tag,
		}
		indexBooksList = append(indexBooksList, res)
	}
	return err, indexBooksList
}
