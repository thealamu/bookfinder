package finder

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/thealamu/bookfinder/internal/pkg/bookdetails"
)

const endpointFmt = "https://www.goodreads.com/search/index.xml?key=%s&q=%s"
const bookLinkFmt = "https://www.goodreads.com/book/isbn/%s"

type author struct {
	Name string `xml:"name"`
}

type bestbook struct {
	ID     string `xml:"id"`
	Title  string `xml:"title"`
	Author author `xml:"author"`
}

type work struct {
	ID       int      `xml:"id"`
	BestBook bestbook `xml:"best_book"`
}

type results struct {
	Work []work `xml:"work"`
}

type search struct {
	Results results `xml:"results"`
}

type goodreadsResponse struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Search  search   `xml:"search"`
}

type GoodReads struct {
	key string
}

func NewGoodReadsFinder(apikey string) *GoodReads {
	return &GoodReads{key: apikey}
}

func (g *GoodReads) Find(s string, c chan bookdetails.BookDetails) {
	var endpointUri = fmt.Sprintf(endpointFmt, g.key, s)
	fmt.Println("goodreads.Find", "GET", endpointUri)
	res, err := http.Get(endpointUri)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	dec := xml.NewDecoder(res.Body)
	var data goodreadsResponse
	err = dec.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Stream bookdetails
	works := data.Search.Results.Work
	for _, w := range works {
		var bookd bookdetails.BookDetails
		bookd.ID = w.BestBook.ID
		bookd.Title = w.BestBook.Title
		bookd.Authors = []string{w.BestBook.Author.Name}
		bookd.Link = fmt.Sprintf(bookLinkFmt, w.BestBook.ID)

		c <- bookd
	}
}
