package manage

import (
	"bookstore/model/common"
)

// MallBooksInfo 结构体
type MallBooksInfo struct {
	BooksId            *int            `json:"booksId" form:"booksId" gorm:"primarykey;AUTO_INCREMENT"`
	BooksName          string          `json:"booksName" form:"booksName" gorm:"column:books_name;comment:图书名;type:varchar(200);"`
	BooksIntro         string          `json:"booksIntro" form:"booksIntro" gorm:"column:books_intro;comment:图书简介;type:varchar(1000);"`
	BooksCategoryId    *int            `json:"booksCategoryId" form:"booksCategoryId" gorm:"column:books_category_id;comment:关联分类id;type:bigint"`
	BooksCoverImg      string          `json:"booksCoverImg" form:"booksCoverImg" gorm:"column:books_cover_img;comment:图书主图;type:varchar(200);"`
	BooksCarousel      string          `json:"booksCarousel" form:"booksCarousel" gorm:"column:books_carousel;comment:图书轮播图;type:varchar(500);"`
	BooksDetailContent string          `json:"booksDetailContent" form:"booksDetailContent" gorm:"column:books_detail_content;comment:图书详情;type:text;"`
	OriginalPrice      *int            `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:图书价格;type:int"`
	SellingPrice       *int            `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:图书实际售价;type:int"`
	StockNum           *int            `json:"stockNum" form:"stockNum" gorm:"column:stock_num;comment:图书库存数量;type:int"`
	Tag                string          `json:"tag" form:"tag" gorm:"column:tag;comment:图书标签;type:varchar(20);"`
	BooksSellStatus    int             `json:"booksSellStatus" form:"booksSellStatus" gorm:"column:books_sell_status;comment:图书上架状态 1-下架 0-上架;type:tinyint"`
	CreateUser         int             `json:"createUser" form:"createUser" gorm:"column:create_user;comment:添加者主键id;type:int"`
	CreateTime         common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:图书添加时间;type:datetime"`
	UpdateUser         int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime         common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:图书修改时间;type:datetime"`
	BooksPublish       string          `json:"booksPublish" form:"booksPublish" gorm:"column:books_publish;comment:出版社出版时间;type:varchar(200);"`
	BooksAuthor        string          `json:"booksAuthor" form:"booksAuthor" gorm:"column:books_author;comment:图书作者;type:varchar(255);"`
	BooksAuthorCountry string          `json:"booksAuthorCountry" form:"booksAuthorCountry" gorm:"column:books_author_country;comment:图书作者国籍;type:varchar(255);"`
}

// TableName MallBooksInfo 表名
func (MallBooksInfo) TableName() string {
	return "tb_bookstore_books_info"
}
