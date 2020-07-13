/* Package find implements book details searching*/
package find

import (
	"fmt"
	"net/http"

	"github.com/thealamu/bookfinder/internal/pkg/server"
)

type FindHandler struct {
	env *server.ServerEnv
}

func NewFindHandler(srvEnv *server.ServerEnv) FindHandler {
	return FindHandler{env: srvEnv}
}

func (f FindHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("find.ServeHTTP", err)
	}
	q := r.Form.Get("q")
	if q == "" {
		fmt.Println("find.ServeHTTP", "Error, no search filter")
		return
	}

	booksList, err := f.env.GoodReadsFinder.Find(q)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, booksList)
}
