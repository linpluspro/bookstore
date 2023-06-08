package response

import (
	"time"
)

type BooksSearchResponse struct {
	BooksId       int    `json:"booksId"`
	BooksName     string `json:"booksName"`
	BooksIntro    string `json:"booksIntro"`
	BooksCoverImg string `json:"booksCoverImg"`
	SellingPrice  int    `json:"sellingPrice"`
}

type BooksInfoDetailResponse struct {
	BooksId                 int                      `json:"booksId"`
	BooksName               string                   `json:"booksName"`
	BooksIntro              string                   `json:"booksIntro"`
	BooksCoverImg           string                   `json:"booksCoverImg"`
	SellingPrice            int                      `json:"sellingPrice"`
	BooksDetailContent      string                   `json:"booksDetailContent"  `
	OriginalPrice           int                      `json:"originalPrice" `
	Tag                     string                   `json:"tag" form:"tag" `
	BooksCarouselList       []string                 `json:"booksCarouselList" `
	BookStoreBookCommentVOS []BookStoreBookCommentVO `json:"comments"`
}

type BookStoreBookCommentVO struct {
	Id          int64          `json:"id"`         //评论ID
	FromId      int64          `json:"fromId"`     //评论者ID
	Name        string         `json:"name"`       //用户姓名
	HeadImg     string         `json:"headImg"`    //用户头像
	Comment     string         `json:"comment"`    //评论内容
	To          string         `json:"to"`         //父评论者姓名
	ToId        int64          `json:"toId"`       //父评论ID 考虑用哪个字段
	Like        int            `json:"like"`       //赞成数量
	CommentNum  int            `json:"commentNum"` //子评论总数
	CommentTime time.Time      `json:"_"`          //评论发布时间
	Time        string         `json:"time"`       //评论发布时间
	BooksId     int64          `json:"booksId"`    //图书ID
	InputShow   bool           `json:"inputShow"`  //是否展示回复输入框
	Reply       []CommentReply `json:"reply"`      // 子回复内容
}

type CommentReply struct {
	Id          int64     `json:"id"`         //评论ID
	FromId      int64     `json:"fromId"`     //评论者ID
	Name        string    `json:"name"`       //用户姓名
	HeadImg     string    `json:"headImg"`    //用户头像
	Comment     string    `json:"comment"`    //评论内容
	To          string    `json:"to"`         //父评论者姓名
	ToId        int64     `json:"toId"`       //父评论ID 考虑用哪个字段
	Like        int       `json:"like"`       //赞成数量
	CommentNum  int       `json:"commentNum"` //子评论总数
	CommentTime time.Time `json:"_"`          //评论发布时间
	Time        string    `json:"time"`       //评论发布时间
	BooksId     int64     `json:"booksId"`    //图书ID
	InputShow   bool      `json:"inputShow"`  //是否展示回复输入框
}
