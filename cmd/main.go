package main

import (
	"github.com/0xff3232/blog/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	userHandler := handler.UserHandler{}

	app.GET("/home", userHandler.HandleUserShow)

	blogHandler := handler.NewBlogHandler()

	app.GET("/article/:path", blogHandler.HandleBlogShow)

	app.Start(":8080")

}
