package manage

import "bookstore/model/common"

type MallOrderItem struct {
	OrderItemId   int             `json:"orderItemId" gorm:"primarykey;AUTO_INCREMENT"`
	OrderId       int             `json:"orderId" form:"orderId" gorm:"column:order_id;;type:bigint"`
	BooksId       int             `json:"booksId" form:"booksId" gorm:"column:books_id;;type:bigint"`
	BooksName     string          `json:"booksName" form:"booksName" gorm:"column:books_name;comment:商品名;type:varchar(200);"`
	BooksCoverImg string          `json:"booksCoverImg" form:"booksCoverImg" gorm:"column:books_cover_img;comment:商品主图;type:varchar(200);"`
	SellingPrice  int             `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品实际售价;type:int"`
	BooksCount    int             `json:"booksCount" form:"booksCount" gorm:"column:books_count;;type:bigint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;type:datetime"`
}

func (MallOrderItem) TableName() string {
	return "tb_bookstore_order_item"
}
