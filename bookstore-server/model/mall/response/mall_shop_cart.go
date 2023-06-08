package response

type CartItemResponse struct {
	CartItemId int `json:"cartItemId"`

	BooksId int `json:"booksId"`

	BooksCount int `json:"booksCount"`

	BooksName string `json:"booksName"`

	BooksCoverImg string `json:"booksCoverImg"`

	SellingPrice int `json:"sellingPrice"`
}
