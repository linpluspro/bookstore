package mall

import (
	"github.com/jinzhu/copier"

	"bookstore/global"
	"bookstore/model/common/enum"
	mallRes "bookstore/model/mall/response"
	"bookstore/model/manage"
)

type MallBooksCategoryService struct {
}

func (m *MallBooksCategoryService) GetCategoriesForIndex() (err error, bookStoreIndexCategoryVOS []mallRes.BookStoreIndexCategoryVO) {
	//获取一级分类的固定数量的数据
	_, firstLevelCategories := selectByLevelAndParentIdsAndNumber([]int{0}, enum.LevelOne.Code(), 10)
	if firstLevelCategories != nil {
		var firstLevelCategoryIds []int
		for _, firstLevelCategory := range firstLevelCategories {
			firstLevelCategoryIds = append(firstLevelCategoryIds, firstLevelCategory.CategoryId)
		}
		//获取二级分类的数据
		_, secondLevelCategories := selectByLevelAndParentIdsAndNumber(firstLevelCategoryIds, enum.LevelTwo.Code(), 0)
		if secondLevelCategories != nil {
			var secondLevelCategoryIds []int
			for _, secondLevelCategory := range secondLevelCategories {
				secondLevelCategoryIds = append(secondLevelCategoryIds, secondLevelCategory.CategoryId)
			}
			//获取三级分类的数据
			_, thirdLevelCategories := selectByLevelAndParentIdsAndNumber(secondLevelCategoryIds, enum.LevelThree.Code(), 0)
			if thirdLevelCategories != nil {
				//根据 parentId 将 thirdLevelCategories 分组
				thirdLevelCategoryMap := make(map[int][]manage.MallBooksCategory)
				for _, thirdLevelCategory := range thirdLevelCategories {
					thirdLevelCategoryMap[thirdLevelCategory.ParentId] = []manage.MallBooksCategory{}
				}
				for k, v := range thirdLevelCategoryMap {
					for _, third := range thirdLevelCategories {
						if k == third.ParentId {
							v = append(v, third)
						}
						thirdLevelCategoryMap[k] = v
					}
				}
				var secondLevelCategoryVOS []mallRes.SecondLevelCategoryVO
				//处理二级分类
				for _, secondLevelCategory := range secondLevelCategories {
					var secondLevelCategoryVO mallRes.SecondLevelCategoryVO
					err = copier.Copy(&secondLevelCategoryVO, &secondLevelCategory)
					//如果该二级分类下有数据则放入 secondLevelCategoryVOS 对象中
					if _, ok := thirdLevelCategoryMap[secondLevelCategory.CategoryId]; ok {
						//根据二级分类的id取出thirdLevelCategoryMap分组中的三级分类list
						tempBooksCategories := thirdLevelCategoryMap[secondLevelCategory.CategoryId]
						var thirdLevelCategoryRes []mallRes.ThirdLevelCategoryVO
						err = copier.Copy(&thirdLevelCategoryRes, &tempBooksCategories)
						secondLevelCategoryVO.ThirdLevelCategoryVOS = thirdLevelCategoryRes
						secondLevelCategoryVOS = append(secondLevelCategoryVOS, secondLevelCategoryVO)
					}

				}
				//处理一级分类
				if secondLevelCategoryVOS != nil {
					//根据 parentId 将 thirdLevelCategories 分组
					secondLevelCategoryVOMap := make(map[int][]mallRes.SecondLevelCategoryVO)
					for _, secondLevelCategory := range secondLevelCategoryVOS {
						secondLevelCategoryVOMap[secondLevelCategory.ParentId] = []mallRes.SecondLevelCategoryVO{}
					}
					for k, v := range secondLevelCategoryVOMap {
						for _, second := range secondLevelCategoryVOS {
							if k == second.ParentId {
								var secondLevelCategory mallRes.SecondLevelCategoryVO
								copier.Copy(&secondLevelCategory, &second)
								v = append(v, secondLevelCategory)
							}
							secondLevelCategoryVOMap[k] = v
						}
					}
					for _, firstCategory := range firstLevelCategories {
						var bookStoreIndexCategoryVO mallRes.BookStoreIndexCategoryVO
						err = copier.Copy(&bookStoreIndexCategoryVO, &firstCategory)
						//如果该一级分类下有数据则放入 bookStoreIndexCategoryVOS 对象中
						if _, ok := secondLevelCategoryVOMap[firstCategory.CategoryId]; ok {
							//根据一级分类的id取出secondLevelCategoryVOMap分组中的二级级分类list
							tempBooksCategories := secondLevelCategoryVOMap[firstCategory.CategoryId]
							bookStoreIndexCategoryVO.SecondLevelCategoryVOS = tempBooksCategories
							bookStoreIndexCategoryVOS = append(bookStoreIndexCategoryVOS, bookStoreIndexCategoryVO)
						}
					}
				}
			}
		}
	}
	return
}

// 获取分类数据
func selectByLevelAndParentIdsAndNumber(ids []int, level int, limit int) (err error, categories []manage.MallBooksCategory) {
	global.GVA_DB.Where("parent_id in ? and category_level =? and is_deleted = 0", ids, level).
		Order("category_rank desc").Limit(limit).Find(&categories)
	return
}
