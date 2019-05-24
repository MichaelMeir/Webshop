package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string
	UUID string
	Products []*Product `gorm:"many2many:product_categories;"`
}
