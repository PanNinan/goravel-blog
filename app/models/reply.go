package models

import "github.com/goravel/framework/database/orm"

type Reply struct {
	orm.Model
	TopicID uint
	UserID  uint
	Content string
}

func (r *Reply) TableName() string {
	return "replies"
}
