package main

import (
	"go-login/db"
	"go-login/handler"
	mdw "go-login/middleware"
	"go-login/models"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := db.OpenDB()
	if err != nil {
		log.Fatal(err.Error())
	}
	db.DB.AutoMigrate(&models.Tokens{}, &models.UserLogin{})

	server := echo.New()
	isLoggedIn := middleware.JWT([]byte("key"))
	isAdmin := mdw.IsAdmin
	server.Use(middleware.Logger())

	server.GET("/", handler.Hello, isLoggedIn)
	server.POST("/login", handler.Login)
	server.GET("/admin", handler.Hello, isLoggedIn, isAdmin)
	server.POST("/signin", handler.SignIn)
	server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))
	// server.Logger.Fatal(server.Start(":3000"))
}
