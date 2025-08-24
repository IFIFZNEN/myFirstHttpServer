package main

import (
	"fmt"
	"net/http"
)

var PORT = ":8081"

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")

}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello", hello)

	server := http.Server{
		Addr:    PORT,
		Handler: router,
	}

	fmt.Printf("Server is listening on port %s\n", PORT)
	server.ListenAndServe()
}
