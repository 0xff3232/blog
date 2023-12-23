package handler

import (
	"github.com/0xff3232/blog/model"
	"github.com/0xff3232/blog/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	// just getting it to work
	u := model.User{
		Name:  "lolollol",
		Email: "test@gmail.com",
	}
	return render(c, user.Show(u))
}
