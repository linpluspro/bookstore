package enum

type BooksStatusEnum int

const (
	BOOKS_DEFAULT BooksStatusEnum = -9
	BOOKS_UNDER   BooksStatusEnum = 0
)

func GetBookStoreBooksStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "已下架"
	default:
		return -9, "error"
	}
}

func (g BooksStatusEnum) Code() int {
	switch g {
	case BOOKS_UNDER:
		return 0
	default:
		return -9
	}
}
