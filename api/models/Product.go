package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name string
	Description string
	UUID string
	Price uint
	Images string
	Categories []*Category `gorm:"many2many:product_categories;"`
}
