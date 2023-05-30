package repository

import (
	"github.com/taka-617/gochat/db"

	"github.com/taka-617/gochat/model"
)

type UserRepository struct{}

func (_ UserRepository) GetAll() ([]model.User, error) {
	dbConn := db.GetDB()
	var u []model.User
	if err := dbConn.Table("users").Select("*").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (_ UserRepository) Create(u model.User) (model.User, error){
	db:= db.GetDB()

	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}