package server

import "net/http"

// ServerEnv defines the environment for a server
type ServerEnv struct {
	Port    string
	Handler http.Handler
}

// NewServerEnv creates and returns a new server environment
//with a default port 8080
func NewServerEnv() *ServerEnv {
	return &ServerEnv{
		Port: "8080",
	}
}
