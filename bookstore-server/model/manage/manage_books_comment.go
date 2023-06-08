package manage

import (
	"time"
)

// MallBooksComment 图书评论表
type MallBooksComment struct {
	Id          int64     `json:"id" gorm:"column:id" db:"id" form:"id"`                                    // 评论ID 时间戳 毫秒
	FromId      int64     `json:"fromId" gorm:"column:from_id" db:"fromId" form:"fromId"`                   //评论者ID
	Name        string    `json:"name" gorm:"column:name" db:"name" form:"name"`                            //用户姓名
	HeadImg     string    `json:"headImg" gorm:"column:head_img" db:"head_img" form:"head_img"`             //用户头像
	Comment     string    `json:"comment" gorm:"column:comment" db:"comment" form:"comment"`                //评论内容
	To          string    `json:"to" gorm:"column:to" db:"to" form:"to"`                                    //父评论者姓名
	ToId        int64     `json:"toId" gorm:"column:to_id" db:"to_id" form:"to_id"`                         //父评论ID 考虑用哪个字段
	Like        int       `json:"like" gorm:"column:like" db:"like" form:"like"`                            //赞成数量
	CommentNum  int       `json:"commentNum" gorm:"column:comment_num" db:"comment_num" form:"comment_num"` //子评论总数
	CommentTime time.Time `json:"time" gorm:"column:comment_time" db:"comment_time" form:"comment_time"`    //评论发布时间
	BooksId     int64     `json:"booksId" gorm:"column:books_id" db:"books_id" form:"books_id"`             //图书ID
	InputShow   bool      `json:"inputShow" gorm:"column:input_show" db:"input_show" form:"input_show"`     //是否展示回复输入框
}

// TableName MallBooksComment 表名
func (MallBooksComment) TableName() string {
	return "tb_bookstore_books_comment"
}
