package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Email string

	Session_ip string
	Session_id string

	Admin bool

	Address string
	ZipCode string
	Province string
	Country string
}
