/* Package find implements book details searching*/
package find

import (
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
	f.env.GoogleBooksFinder.Find("Stuff")
}
