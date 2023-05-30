package db

import (
	"log"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {
	dsn := "root:mysql@tcp(mysql:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}
}

// GetDB is called in models
func GetDB() (*gorm.DB) {
	return db
}
