package finder

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thealamu/bookfinder/internal/pkg/bookdetails"
)

const EndpointFmt = "https://www.googleapis.com/books/v1/volumes?q=%s"

type volumeInfo struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Authors     []string `json:"authors"`
	Description string   `json:"description"`
	ReadLink    string   `json:"previewLink"`
}

type volume struct {
	ID         string     `json:"id"`
	VolumeInfo volumeInfo `json:"volumeInfo"`
}

type response struct {
	TotalItems int      `json:"totalItems"`
	Items      []volume `json:"items"`
}

type GoogleBooks struct{}

func NewGoogleBooksFinder() *GoogleBooks {
	return &GoogleBooks{}
}

func (g *GoogleBooks) Find(s string) ([]bookdetails.BookDetails, error) {
	endpointUri := fmt.Sprintf(EndpointFmt, s)
	fmt.Println("googlebooks.Find", "GET", endpointUri)

	// Go get results from googleapis
	r, err := http.Get(endpointUri)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var res response
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	//return the result
	var results []bookdetails.BookDetails
	for i := 0; i < len(res.Items); i++ {
		bd := bookdetails.BookDetails{}
		item := res.Items[i]
		vInfo := item.VolumeInfo

		bd.ID = item.ID
		bd.Title = vInfo.Title
		bd.Subtitle = vInfo.Subtitle
		bd.Authors = vInfo.Authors
		bd.Desc = vInfo.Description
		bd.Link = vInfo.ReadLink

		results = append(results, bd)
	}

	return results, nil
}
