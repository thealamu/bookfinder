package server

import (
	"context"
	"fmt"
	"net/http"
)

// StartServer starts the server using the environment
func StartServer(ctx context.Context, env *ServerEnv) {
	addr := ":" + env.Port
	fmt.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, env.Handler)
	if err != nil {
		panic(err)
	}
}
