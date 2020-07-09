package server

import (
	"context"
	"fmt"
)

// StartServer starts the server using the environment
func StartServer(ctx context.Context, env *ServerEnv) {
	fmt.Println("Starting server on port", env.Port)
}
