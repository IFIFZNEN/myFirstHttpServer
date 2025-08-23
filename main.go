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
	http.HandleFunc("/hello", hello)
	fmt.Printf("Server is listening on port %s\n", PORT)
	http.ListenAndServe(PORT, nil)

}
