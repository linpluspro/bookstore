package enum

type IndexConfigEnum int8

// 首页配置项 1-搜索框热搜 2-搜索下拉框热搜 3-(首页)热销商品 4-(首页)新品上线 5-(首页)为你推荐
const (
	IndexSearchHots     IndexConfigEnum = 1
	IndexSearchDownHots IndexConfigEnum = 2
	IndexBooksHot       IndexConfigEnum = 3
	IndexBooksNew       IndexConfigEnum = 4
	IndexBooksRecommond IndexConfigEnum = 5
)

func (i IndexConfigEnum) Info() (int, string) {
	switch i {
	case IndexSearchHots:
		return 1, "INDEX_SEARCH_HOTS"
	case IndexSearchDownHots:
		return 2, "二级分类"
	case IndexBooksHot:
		return 3, "三级分类"
	case IndexBooksNew:
		return 4, "三级分类"
	case IndexBooksRecommond:
		return 5, "三级分类"
	default:
		return 0, "DEFAULT"
	}
}

func (i IndexConfigEnum) Code() int {
	switch i {
	case IndexSearchHots:
		return 1
	case IndexSearchDownHots:
		return 2
	case IndexBooksHot:
		return 3
	case IndexBooksNew:
		return 4
	case IndexBooksRecommond:
		return 5
	default:
		return 0
	}
}
