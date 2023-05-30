package repository

import (
	"github.com/taka-617/gochat/db"

	"github.com/taka-617/gochat/model"
)

type OauthRepository struct{}

func (_ OauthRepository) Create(o model.Oauth) (model.Oauth, error){
	db:= db.GetDB()

	if err := db.Create(&o).Error; err != nil {
		return o, err
	}
	return o, nil
}

func (_ OauthRepository) GetFirst(o model.Oauth) (model.Oauth, error) {
	dbConn := db.GetDB()
	oauth := model.Oauth{}

	if err := dbConn.Where(&o).First(&oauth).Error; err != nil {
		return oauth, err
	}

	return oauth, nil
}