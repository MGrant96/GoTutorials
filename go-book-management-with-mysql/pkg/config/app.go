package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	// Open Connection with DB
	d, err := gorm.Open("mysql", "admin:admin@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=true&loc=Local") // Username, Password, Name of db table, charset=utf8&parseTime=true&loc=local needs to be in it for MySql
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}