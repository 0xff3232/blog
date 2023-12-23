package model

import (
	"bytes"
	"log"
	"os"

	"github.com/0xff3232/blog/types"
	"github.com/yuin/goldmark"
)

type PostModeler interface {
	MarkdownToHTML(string) (string error) // will take a .md and convert to html
}

type PostModel struct {
	post types.Post
}

func (s *PostModel) MarkdownToHTML(path string) (string, error) {

	markdown, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Failed to read markdown file: %v", err)
		return "", err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		log.Printf("Failed to convert markdown to HTML: %v", err)
		return "", err
	}

	return buf.String(), nil

}
