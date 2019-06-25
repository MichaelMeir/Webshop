package models

import (
	"github.com/jinzhu/gorm"
	"webshop/Env"
)
import _ "github.com/go-sql-driver/mysql"


type Database struct {
	Data *gorm.DB
}

func OpenDatabase() *Database {
	db, err := gorm.Open("mysql", Env.LoadENV()["DB_USER"] + ":" + Env.LoadENV()["DB_PASS"] + "@/" + Env.LoadENV()["DB"] + "?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	var database = &Database{db}
	database.DbInit()
	return database
}

func(db *Database) Close() {
	db.Data.Close()
}

func(db *Database) DbInit() {
	db.Data.AutoMigrate(Product{}, Category{}, Image{}, User{}, Order{})
	db.Data.Model(&Product{}).Related(&[]Category{}, "Categories")
}