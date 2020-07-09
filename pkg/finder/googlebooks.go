package finder

import (
	"errors"
	"fmt"

	"github.com/thealamu/bookfinder/internal/pkg/bookdetails"
)

type GoogleBooks struct {
	apiKey string
}

func NewGoogleBooksFinder(key string) (*GoogleBooks, error) {
	if key == "" {
		fmt.Println("finder.NewGoogleBooksFinder", "No API key")
		return nil, errors.New("No Google Books API key")
	}

	return &GoogleBooks{apiKey: key}, nil
}

func (g *GoogleBooks) Find(s string) *bookdetails.BookDetails {
	return nil
}
