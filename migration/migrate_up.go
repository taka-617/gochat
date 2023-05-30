package main

import (
	"github.com/taka-617/gochat/db"
	"github.com/taka-617/gochat/model"
)

func migrateUp() {
	db := db.GetDB()

	db.AutoMigrate(&model.User{}, &model.Oauth{})
}