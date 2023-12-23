package handler

import (
	"net/http"

	"github.com/0xff3232/blog/model"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	postModel model.PostModel
}

func NewPostHanlder(postModel model.PostModel) *PostHandler {
	return &PostHandler{
		postModel: postModel,
	}
}

func (h *PostHandler) HandleCreatePost(c echo.Context, filePath string) error {
	htmlContent, err := h.postModel.MarkdownToHTML("markdown/" + filePath)
	if err != nil {
		// Handle the error appropriately
		return err
	}

	return c.HTML(http.StatusOK, htmlContent)
}
