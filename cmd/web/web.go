package main

import (
	"net/http"
)

func server() {
	server := http.NewServeMux()
	port := "4000"
	server.Handle("/", http.FileServer(http.Dir("static/build")))
	http.ListenAndServe(":"+port, server)
}

func main() {

	// Start server.
	server()
}
