package main

import (
	"context"
	"flag"

	"github.com/gorilla/mux"
	"github.com/thealamu/bookfinder/internal/pkg/find"
	"github.com/thealamu/bookfinder/internal/pkg/home"
	"github.com/thealamu/bookfinder/internal/pkg/server"
	"github.com/thealamu/bookfinder/internal/pkg/setup"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "port to start the server on")
}

func main() {
	ctx := context.Background()

	flag.Parse()

	srvEnv := setup.ServerEnv(port)

	router := mux.NewRouter()
	router.HandleFunc("/", home.HomeHandler)
	router.Handle("/api/find", find.NewFindHandler())
	srvEnv.Handler = router

	server.StartServer(ctx, srvEnv)
}
