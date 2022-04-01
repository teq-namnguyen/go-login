package db

import (
	"errors"
	"go-login/models"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDB() error {
	// dns := "host=localhost user=postgres dbname=login_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func AddToken(token string, time time.Time) error {
	result := DB.Create(&models.Tokens{Token: token, Exp: time})
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return nil
}

func AddUserNamePass(user *models.UserLogin) error {
	result := DB.Create(user)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return nil
}

func CheckUserIsNotExist(username string) error {
	var user models.UserLogin
	result := DB.First(&user, "username = ?", username)
	log.Printf("%#v", result)
	if result.Error != nil {
		return nil
	}
	return errors.New("error")
}

func GetUserByUserName(username string) (*models.UserLogin, error) {
	var user models.UserLogin

	result := DB.Find(&user, "username = ?", username)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return &user, nil
}
