package main

import (
	"fmt"
	"net/http"

	"github.com/rest/client"
	"github.com/rest/server"
)

func main() {
	mux := http.NewServeMux()

	// Client:
	mux.HandleFunc("/page/student", client.Student)

	// Server:
	mux.HandleFunc("/api/student", server.Students)

	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", mux)
}
