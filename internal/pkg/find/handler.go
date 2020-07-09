/* Package find implements book details searching*/
package find

import (
	"fmt"
	"net/http"
)

type FindHandler struct{}

func NewFindHandler() FindHandler {
	return FindHandler{}
}

func (f FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Search here!")
}
