package main

import (
	"context"
	"flag"
	"fmt"

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

	srvEnv, err := setup.ServerEnv(port)
	if err != nil {
		fmt.Println("main.Main", "Could not setup server env")
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", home.HomeHandler)
	router.Handle("/api/find", find.NewFindHandler(srvEnv))
	srvEnv.Handler = router

	if err := server.StartServer(ctx, srvEnv); err != nil {
		fmt.Println("main.Main", "Could not start server")
		panic(err)
	}
}
