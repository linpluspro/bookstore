package request

import (
	"bookstore/model/common"
	"bookstore/model/common/request"
	"bookstore/model/manage"
)

type MallBooksInfoSearch struct {
	manage.MallBooksInfo
	request.PageInfo
}

type BooksInfoAddParam struct {
	BooksName          string `json:"booksName"`
	BooksPublish       string `json:"booksPublish"`
	BooksAuthor        string `json:"booksAuthor"`
	BooksAuthorCountry string `json:"booksAuthorCountry"`
	BooksIntro         string `json:"booksIntro"`
	BooksCategoryId    int    `json:"booksCategoryId"`
	BooksCoverImg      string `json:"booksCoverImg"`
	BooksCarousel      string `json:"booksCarousel"`
	BooksDetailContent string `json:"booksDetailContent"`
	OriginalPrice      int    `json:"originalPrice"`
	SellingPrice       int    `json:"sellingPrice"`
	StockNum           int    `json:"stockNum"`
	Tag                string `json:"tag"`
	BooksSellStatus    int    `json:"booksSellStatus"`
}

// BooksInfoUpdateParam 更新商品信息的入参
type BooksInfoUpdateParam struct {
	BooksId            string          `json:"booksId"`
	BooksName          string          `json:"booksName"`
	BooksPublish       string          `json:"booksPublish"`
	BooksAuthor        string          `json:"booksAuthor"`
	BooksAuthorCountry string          `json:"booksAuthorCountry"`
	BooksIntro         string          `json:"booksIntro"`
	BooksCategoryId    int             `json:"booksCategoryId"`
	BooksCoverImg      string          `json:"booksCoverImg"`
	BooksCarousel      string          `json:"booksCarousel"`
	BooksDetailContent string          `json:"booksDetailContent"`
	OriginalPrice      int             `json:"originalPrice"`
	SellingPrice       int             `json:"sellingPrice"`
	StockNum           int             `json:"stockNum"`
	Tag                string          `json:"tag"`
	BooksSellStatus    int             `json:"booksSellStatus"`
	UpdateUser         int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime         common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改时间;type:datetime"`
}

type StockNumDTO struct {
	BooksId    int `json:"booksId"`
	BooksCount int `json:"booksCount"`
}
