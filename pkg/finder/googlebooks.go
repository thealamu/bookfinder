package finder

import (
	"fmt"
	"net/http"

	"github.com/thealamu/bookfinder/internal/pkg/bookdetails"
)

const EndpointFmt = "https://www.googleapis.com/books/v1/volumes?q=%s"

type GoogleBooks struct{}

func NewGoogleBooksFinder() *GoogleBooks {
	return &GoogleBooks{}
}

func (g *GoogleBooks) Find(s string) (*bookdetails.BookDetails, error) {
	endpointUri := fmt.Sprintf(EndpointFmt, s)
	fmt.Println(endpointUri)

	// Go get results from googleapis
	resp, err := http.Get(endpointUri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return nil, nil
}
