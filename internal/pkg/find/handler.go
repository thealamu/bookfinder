/* Package find implements book details searching*/
package find

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/thealamu/bookfinder/internal/pkg/bookdetails"
	"github.com/thealamu/bookfinder/internal/pkg/server"
)

type response struct {
	Query  string                    `json:"query"`
	Result []bookdetails.BookDetails `json:"result"`
}

type FindHandler struct {
	env *server.ServerEnv
}

func NewFindHandler(srvEnv *server.ServerEnv) FindHandler {
	return FindHandler{env: srvEnv}
}

func (f FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("find.ServeHTTP", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Could not parse queries")
	}
	q := r.Form.Get("q")
	if q == "" {
		fmt.Println("find.ServeHTTP", "Error, no search filter")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "No search filter")
		return
	}
	//Escape querystring
	q = url.QueryEscape(q)

	// Create result channels
	allChan := make(chan bookdetails.BookDetails, 40)
	n := 2

	//Goodreads
	go func() {
		f.env.GoodReadsFinder.Find(q, allChan)
		n--
		if n <= 0 {
			close(allChan)
		}
	}()
	//Google Books
	go func() {
		f.env.GoogleBooksFinder.Find(q, allChan)
		n--
		if n <= 0 {
			close(allChan)
		}
	}()

	// Populate data with channel data
	var booksList []bookdetails.BookDetails
	for book := range allChan {
		booksList = append(booksList, book)
	}

	var resp response
	resp.Query = q
	resp.Result = booksList

	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("find.Handler serving %d books\n", len(booksList))
	fmt.Fprintln(w, string(data))
}
