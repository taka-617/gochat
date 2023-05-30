package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/taka-617/gochat/model"
	"github.com/taka-617/gochat/repository"

	"gorm.io/gorm"
)

type AuthService struct{}

func (_ AuthService) LoginOauth(provider string, id string, name string, email string) {
	var oauth_repository repository.OauthRepository
	var user_repository repository.UserRepository

	o := model.Oauth{Name: provider, SnsId: id}
	oauth, err := oauth_repository.GetFirst(o)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// oauthの情報がない場合はユーザー作成を行う
		u := model.User{Name: name, Email: email}

		uc, err := user_repository.Create(u)
		if err != nil {
			log.Fatalln("create error")
		}

		oc := model.Oauth{Name: provider, SnsId: id, UserID: int(uc.ID)}
		_, err = oauth_repository.Create(oc)
		if err != nil {
			log.Fatalln("create error")
		}
	}

	fmt.Println(oauth)
}
