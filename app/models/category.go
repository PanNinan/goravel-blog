package models

type Category struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	Name        string
	Description string
	PostCount   uint
}

func (r *Category) TableName() string {
	return "categories"
}
