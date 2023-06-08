package response

import "bookstore/model/manage"

type BooksCategoryResponse struct {
	BooksCategory manage.MallBooksCategory `json:"mallBooksCategory"`
}
