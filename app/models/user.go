package models

import (
	"github.com/goravel/framework/database/orm"
)

type User struct {
	orm.Model
	Name         string
	Email        string
	Password     string
	Introduction string
	Avatar       string
}

func (r *User) TableName() string {
	return "users"
}
