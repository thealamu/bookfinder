package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// StartServer starts the server using the environment
func StartServer(ctx context.Context, env *ServerEnv) {
	addr := net.JoinHostPort("localhost", env.Port)
	fmt.Println("Starting server on", addr)
	err := http.ListenAndServe(addr, env.Handler)
	if err != nil {
		panic(err)
	}
}
