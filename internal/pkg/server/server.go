package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// StartServer starts the server using the environment
func StartServer(ctx context.Context, env *ServerEnv) error {
	addr := net.JoinHostPort("localhost", env.Port)
	fmt.Println("Starting server on", addr)
	return http.ListenAndServe(addr, env.Handler)
}
