package components

import (
	"bytes"
	"context"
	"io"
	"os"

	"github.com/yuin/goldmark"
)

type MarkdownComponent struct {
	FilePath string
}

// Standard method definition with the receiver (mc MarkdownComponent)
func (mc MarkdownComponent) Render(ctx context.Context, w io.Writer) error {
	markdown, err := os.ReadFile(mc.FilePath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		return err
	}

	_, err = w.Write(buf.Bytes())
	return err
}
