package models

import (
	"github.com/goravel/framework/database/orm"
	"github.com/goravel/framework/support/carbon"
)

type Notification struct {
	orm.Timestamps
	Id             string `gorm:"primaryKey" json:"id"`
	Type           string
	NotifiableType string
	NotifiableId   uint
	Data           string
	ReadAt         *carbon.DateTime
}

func (r *Notification) TableName() string {
	return "notifications"
}
