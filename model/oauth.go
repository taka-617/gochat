package model

import (
	"gorm.io/gorm"
)

type Oauth struct {
	gorm.Model
	Name      string `gorm:"column:name"`
	SnsId     string `gorm:"column:sns_id"`
	UserID    int `gorm:"column:user_id;"`
	User      User
}