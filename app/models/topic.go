package models

import "github.com/goravel/framework/database/orm"

type Topic struct {
	orm.Model
	CategoryID      uint
	UserID          uint
	Title           string
	Content         string
	ReplyCount      uint
	ViewCount       uint
	LastReplyUserID uint
	Order           uint
	Excerpt         string
	Slug            string
}

func (r *Topic) TableName() string {
	return "topics"
}
