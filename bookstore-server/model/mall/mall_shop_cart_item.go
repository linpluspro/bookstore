package mall

import (
	"bookstore/model/common"
)

// MallShoppingCartItem 结构体
// 如果含有time.Time 请自行import time包
type MallShoppingCartItem struct {
	CartItemId int             `json:"cartItemId" form:"cartItemId" gorm:"primarykey;AUTO_INCREMENT"`
	UserId     int             `json:"userId" form:"userId" gorm:"column:user_id;comment:用户主键id;type:bigint"`
	BooksId    int             `json:"booksId" form:"booksId" gorm:"column:books_id;comment:关联商品id;type:bigint"`
	BooksCount int             `json:"booksCount" form:"booksCount" gorm:"column:books_count;comment:数量(最大为5);type:int"`
	IsDeleted  int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:删除标识字段(0-未删除 1-已删除);type:tinyint"`
	CreateTime common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
	UpdateTime common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:最新修改时间;type:datetime"`
}

// TableName MallShoppingCartItem 表名
func (MallShoppingCartItem) TableName() string {
	return "tb_bookstore_shopping_cart_item"
}
