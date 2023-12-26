package handler

import (
	"bytes"
	"os"

	"github.com/0xff3232/blog/view/blogpost"
	"github.com/labstack/echo/v4"
	"github.com/yuin/goldmark"
)

type BlogHandler struct {
}

func NewBlogHandler() *BlogHandler {
	return &BlogHandler{
		// Initialize dependencies
	}
}

func (h BlogHandler) HandleBlogShow(c echo.Context) error {
	path := c.Param("path")
	filePath := "markdown/" + path

	// Read and convert Markdown to HTML
	markdown, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		return err
	}
	htmlContent := buf.String()

	// Pass the HTML content to the template
	return blogpost.BlogPost(htmlContent).Render(c.Request().Context(), c.Response().Writer)
}
