package manage

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"bookstore/global"
	"bookstore/model/common"
	"bookstore/model/common/request"
	"bookstore/model/manage"
	manageReq "bookstore/model/manage/request"
	"bookstore/utils"
)

type ManageBooksCategoryService struct {
}

// AddCategory 添加分类
func (m *ManageBooksCategoryService) AddCategory(req manageReq.MallBooksCategoryReq) (err error) {
	if !errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		req.CategoryLevel, req.CategoryName).First(&manage.MallBooksCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同分类")
	}

	category := manage.MallBooksCategory{
		CategoryLevel: req.CategoryLevel,
		CategoryName:  req.CategoryName,
		CategoryRank:  req.CategoryRank,
		ParentId:      req.ParentId,
		IsDeleted:     0,
		CreateTime:    common.JSONTime{Time: time.Now()},
		UpdateTime:    common.JSONTime{Time: time.Now()},
	}
	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err = utils.Verify(category, utils.BooksCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Create(&category).Error
}

// UpdateCategory 更新商品分类
func (m *ManageBooksCategoryService) UpdateCategory(req manageReq.MallBooksCategoryReq) (err error) {
	if errors.Is(global.GVA_DB.Where("category_level=? AND category_name=? AND is_deleted=0",
		req.CategoryLevel, req.CategoryName).First(&manage.MallBooksCategory{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同分类")
	}

	category := manage.MallBooksCategory{
		CategoryName: req.CategoryName,
		CategoryRank: req.CategoryRank,
		UpdateTime:   common.JSONTime{Time: time.Now()},
	}
	// 这个校验理论上应该放在api层，但是因为前端的传值是string，而我们的校验规则是Int,所以只能转换格式后再校验
	if err := utils.Verify(category, utils.BooksCategoryVerify); err != nil {
		return errors.New(err.Error())
	}
	return global.GVA_DB.Where("category_id =?", req.CategoryId).Updates(&category).Error

}

// SelectCategoryPage 获取分类分页数据
func (m *ManageBooksCategoryService) SelectCategoryPage(req manageReq.SearchCategoryParams) (err error, list interface{}, total int64) {
	limit := req.PageSize
	if limit > 1000 {
		limit = 1000
	}
	offset := req.PageSize * (req.PageNumber - 1)
	db := global.GVA_DB.Model(&manage.MallBooksCategory{})
	var categoryList []manage.MallBooksCategory

	if utils.NumsInList(req.CategoryLevel, []int{1, 2, 3}) {
		db.Where("category_level=?", req.CategoryLevel)
	}
	if req.ParentId >= 0 {
		db.Where("parent_id=?", req.ParentId)
	}
	err = db.Where("is_deleted=0").Count(&total).Error

	if err != nil {
		return err, categoryList, total

	} else {
		db = db.Where("is_deleted=0").Order("category_rank desc").Limit(limit).Offset(offset)
		err = db.Find(&categoryList).Error
	}
	return err, categoryList, total
}

// SelectCategoryById 获取单个分类数据
func (m *ManageBooksCategoryService) SelectCategoryById(categoryId int) (err error, booksCategory manage.MallBooksCategory) {
	err = global.GVA_DB.Where("category_id=?", categoryId).First(&booksCategory).Error
	return err, booksCategory
}

// DeleteBooksCategoriesByIds 批量设置失效
func (m *ManageBooksCategoryService) DeleteBooksCategoriesByIds(ids request.IdsReq) (err error, booksCategory manage.MallBooksCategory) {
	err = global.GVA_DB.Where("category_id in ?", ids.Ids).UpdateColumns(manage.MallBooksCategory{IsDeleted: 1}).Error
	return err, booksCategory
}

func (m *ManageBooksCategoryService) SelectByLevelAndParentIdsAndNumber(parentId int, categoryLevel int) (err error, booksCategories []manage.MallBooksCategory) {
	err = global.GVA_DB.Where("category_id in ?", parentId).Where("category_level=?", categoryLevel).Where("is_deleted=0").Error
	return err, booksCategories

}
