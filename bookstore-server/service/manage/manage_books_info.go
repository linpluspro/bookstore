package manage

import (
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"

	"bookstore/global"
	"bookstore/model/common"
	"bookstore/model/common/enum"
	"bookstore/model/common/request"
	"bookstore/model/manage"
	manageReq "bookstore/model/manage/request"
)

type ManageBooksInfoService struct {
}

// CreateMallBooksInfo 创建MallBooksInfo
func (m *ManageBooksInfoService) CreateMallBooksInfo(req manageReq.BooksInfoAddParam) (err error) {
	var booksCategory manage.MallBooksCategory
	err = global.GVA_DB.Where("category_id=?  AND is_deleted=0", req.BooksCategoryId).First(&booksCategory).Error
	if booksCategory.CategoryLevel != enum.LevelThree.Code() {
		return errors.New("分类数据异常")
	}
	if !errors.Is(global.GVA_DB.Where("books_name=? AND books_category_id=?", req.BooksName, req.BooksCategoryId).First(&manage.MallBooksInfo{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("已存在相同的商品信息")
	}

	booksInfo := manage.MallBooksInfo{
		BooksName:          req.BooksName,
		BooksAuthor:        req.BooksAuthor,
		BooksAuthorCountry: req.BooksAuthorCountry,
		BooksPublish:       req.BooksPublish,
		BooksIntro:         req.BooksIntro,
		BooksCategoryId:    &req.BooksCategoryId,
		BooksCoverImg:      req.BooksCoverImg,
		BooksDetailContent: req.BooksDetailContent,
		OriginalPrice:      &req.OriginalPrice,
		SellingPrice:       &req.SellingPrice,
		StockNum:           &req.StockNum,
		Tag:                req.Tag,
		BooksSellStatus:    req.BooksSellStatus,
		CreateTime:         common.JSONTime{Time: time.Now()},
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}

	err = global.GVA_DB.Create(&booksInfo).Error
	return err
}

// DeleteMallBooksInfo 删除MallBooksInfo记录
func (m *ManageBooksInfoService) DeleteMallBooksInfo(mallBooksInfo manage.MallBooksInfo) (err error) {
	err = global.GVA_DB.Delete(&mallBooksInfo).Error
	return err
}

// ChangeMallBooksInfoByIds 上下架
func (m *ManageBooksInfoService) ChangeMallBooksInfoByIds(ids request.IdsReq, sellStatus string) (err error) {
	intSellStatus, _ := strconv.Atoi(sellStatus)
	//更新字段为0时，不能直接UpdateColumns
	err = global.GVA_DB.Model(&manage.MallBooksInfo{}).Where("books_id in ?", ids.Ids).Update("books_sell_status", intSellStatus).Error
	return err
}

// UpdateMallBooksInfo 更新MallBooksInfo记录
func (m *ManageBooksInfoService) UpdateMallBooksInfo(req manageReq.BooksInfoUpdateParam) (err error) {
	booksId, _ := strconv.Atoi(req.BooksId)
	booksInfo := manage.MallBooksInfo{
		BooksId:            &booksId,
		BooksName:          req.BooksName,
		BooksAuthor:        req.BooksAuthor,
		BooksAuthorCountry: req.BooksAuthorCountry,
		BooksPublish:       req.BooksPublish,
		BooksIntro:         req.BooksIntro,
		BooksCategoryId:    &req.BooksCategoryId,
		BooksCoverImg:      req.BooksCoverImg,
		BooksDetailContent: req.BooksDetailContent,
		OriginalPrice:      &req.OriginalPrice,
		SellingPrice:       &req.SellingPrice,
		StockNum:           &req.StockNum,
		Tag:                req.Tag,
		BooksSellStatus:    req.BooksSellStatus,
		UpdateTime:         common.JSONTime{Time: time.Now()},
	}

	err = global.GVA_DB.Where("books_id=?", booksInfo.BooksId).Updates(&booksInfo).Error
	return err
}

// GetMallBooksInfo 根据id获取MallBooksInfo记录
func (m *ManageBooksInfoService) GetMallBooksInfo(id int) (err error, mallBooksInfo manage.MallBooksInfo) {
	err = global.GVA_DB.Where("books_id = ?", id).First(&mallBooksInfo).Error
	return
}

// GetMallBooksInfoInfoList 分页获取MallBooksInfo记录
func (m *ManageBooksInfoService) GetMallBooksInfoInfoList(info manageReq.MallBooksInfoSearch, booksName string, booksSellStatus int) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallBooksInfo{})
	var mallBooksInfos []manage.MallBooksInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if booksName != "" {
		db.Where("books_name =?", booksName)
	}

	db.Where("books_sell_status =?", booksSellStatus)

	err = db.Limit(limit).Offset(offset).Order("books_id desc").Find(&mallBooksInfos).Error
	return err, mallBooksInfos, total
}
