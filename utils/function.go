package utils

import (
	"go-login/models"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswordBscrypt(user *models.UserLogin) error {
	passwordByte := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	user.Password = string(hashedPassword)
	return nil
}
