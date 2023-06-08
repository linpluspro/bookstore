package request

import (
	"bookstore/model/common/request"
)

type MallShopCartSearch struct {
	request.PageInfo
}

type SaveCartItemParam struct {
	BooksCount int `json:"booksCount"`
	BooksId    int `json:"booksId"`
}

type UpdateCartItemParam struct {
	CartItemId int `json:"cartItemId"`
	BooksCount int `json:"booksCount"`
}
