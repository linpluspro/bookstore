package response

type MallIndexConfigBooksResponse struct {
	BooksId       int    `json:"booksId"`
	BooksName     string `json:"booksName"`
	BooksIntro    string `json:"booksIntro"`
	BooksCoverImg string `json:"booksCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
	Tag           string `json:"tag"`
}
