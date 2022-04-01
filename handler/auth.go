package handler

import (
	"go-login/db"
	"go-login/models"
	"go-login/utils"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	u := new(models.UserLogin)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := db.GetUserByUserName(u.Username)
	if err != nil {
		log.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		log.Fatal(err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	timeExp := time.Now().Add(2 * time.Hour)
	claims["username"] = u.Username
	claims["admin"] = false
	claims["exp"] = timeExp.Unix()
	t, err := token.SignedString([]byte("key"))
	if err != nil {
		log.Printf("signed token err %v\n", err)
		return err
	}
	db.AddToken(t, timeExp.UTC())
	return c.JSON(http.StatusOK, &models.Tokens{
		Token: t,
		Exp:   time.Now().Add(2 * time.Hour).UTC(),
	})
}

func SignIn(c echo.Context) error {
	u := new(models.UserLogin)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := db.CheckUserIsNotExist(u.Username)
	log.Print(err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, u.Username)
	}
	err = utils.HashPasswordBscrypt(u)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AddUserNamePass(u)
	if err != nil {
		log.Fatalf("%#v", err.Error())
	}
	return c.JSON(http.StatusOK, u.Username)
}
