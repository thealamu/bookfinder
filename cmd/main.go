package main

import (
	"context"
	"flag"

	"github.com/gorilla/mux"
	"github.com/thealamu/bookfinder/internal/pkg/server"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "port to start the server on")
}

func main() {
	ctx := context.Background()

	flag.Parse()

	srvEnv := server.NewServerEnv()
	srvEnv.Port = port
	router := mux.NewRouter()
	srvEnv.Handler = router

	server.StartServer(ctx, srvEnv)
}
