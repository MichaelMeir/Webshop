package models

import "github.com/jinzhu/gorm"

type Order struct {
	gorm.Model
	Products string `sql:"type:longtext;"`
	Price uint
	User string `sql:"type:longtext;"`
}
