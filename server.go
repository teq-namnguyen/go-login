package main

import (
	"go-login/handler"
	mdw "go-login/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()
	isLoggedIn := middleware.JWT([]byte("key"))
	isAdmin := mdw.IsAdmin
	server.Use(middleware.Logger())

	server.GET("/", handler.Hello, isLoggedIn)
	server.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))
	server.GET("/admin", handler.Hello, isLoggedIn, isAdmin)
	server.Logger.Fatal(server.Start(":80"))
}
