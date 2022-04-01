package models

import (
	"time"

	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	Token string    `json:"token"`
	Exp   time.Time `json:"exp"`
}

type X struct {
	Text string `json:"text"`
}

type UserLogin struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique"`
	Password string `json:"password" form:"password"`
}
