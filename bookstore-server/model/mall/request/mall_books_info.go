package request

import (
	"time"
)

type BooksSearchParams struct {
	Keyword         string `form:"keyword"`
	BooksCategoryId int    `form:"booksCategoryId"`
	OrderBy         string `form:"orderBy"`
	PageNumber      int    `form:"pageNumber"`
}

// CreateBookCommentParam 图书评论信息
type CreateBookCommentParam struct {
	BooksID     int64     `json:"booksId"`   //图书ID
	Comment     string    `json:"comment"`   //评论内容
	To          string    `json:"to"`        //父评论者姓名
	ToId        int64     `json:"toId"`      //父评论ID
	CommentTime time.Time `json:"time"`      //评论发布时间
	InputShow   bool      `json:"inputShow"` //是否展示回复输入框
	FromId      int64     `json:"fromId"`    //评论者ID
	Name        string    `json:"name"`      //用户姓名
}
