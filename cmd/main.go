package main

import (
	"fmt"
	"net/http"

	"myFirstHttpServer/configs"
	"myFirstHttpServer/internal/auth"
	"myFirstHttpServer/pkg/db"
)

var PORT = ":8081"

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDB(conf)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Server is listening on port %s\n", PORT)
	server.ListenAndServe()
}
