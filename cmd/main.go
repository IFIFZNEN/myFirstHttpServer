package main

import (
	"fmt"
	"net/http"

	"myFirstHttpServer/configs"
	"myFirstHttpServer/internal/auth"
	"myFirstHttpServer/internal/link"
	"myFirstHttpServer/pkg/db"
)

var PORT = ":8081"

func main() {
	conf := configs.LoadConfig()
	db := db.NewDB(conf)
	router := http.NewServeMux()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	// Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewHandlerLink(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})
	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Server is listening on port %s\n", PORT)
	server.ListenAndServe()
}
