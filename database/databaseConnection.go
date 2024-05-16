package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/tufind-db?charset=utf8&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = database
}
