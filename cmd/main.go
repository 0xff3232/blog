package main

import (
	"github.com/0xff3232/blog/handler"
	"github.com/0xff3232/blog/model"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	postModel := model.PostModel{}

	postHandler := handler.NewPostHanlder(postModel)
	userHandler := handler.UserHandler{}

	app.GET("/home", userHandler.HandleUserShow)

	app.GET("/test", func(c echo.Context) error {
		return postHandler.HandleCreatePost(c, "test.md")
	})

	app.Start(":8080")

}
