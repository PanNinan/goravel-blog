package models

import "github.com/goravel/framework/database/orm"

type Link struct {
	orm.Model
	Title string
	Link  string
}

func (r *Link) TableName() string {
	return "links"
}
