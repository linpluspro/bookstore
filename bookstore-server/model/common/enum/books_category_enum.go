package enum

type BooksCategoryLevel int8

const (
	Default    BooksCategoryLevel = 0
	LevelOne   BooksCategoryLevel = 1
	LevelTwo   BooksCategoryLevel = 2
	LevelThree BooksCategoryLevel = 3
)

func (g BooksCategoryLevel) Info() (int, string) {
	switch g {
	case LevelOne:
		return 1, "一级分类"
	case LevelTwo:
		return 2, "二级分类"
	case LevelThree:
		return 3, "三级分类"
	default:
		return 0, "error"
	}
}

func (g BooksCategoryLevel) Code() int {
	switch g {
	case LevelOne:
		return 1
	case LevelTwo:
		return 2
	case LevelThree:
		return 3
	default:
		return 0
	}
}
