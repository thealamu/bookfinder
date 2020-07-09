package finder

import "github.com/thealamu/bookfinder/internal/pkg/bookdetails"

type GoogleBooks struct{}

func NewGoogleBooksFinder() *GoogleBooks {
	return &GoogleBooks{}
}

func (g *GoogleBooks) Find(s string) *bookdetails.BookDetails {
	return nil
}
