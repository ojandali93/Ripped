package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define your server configuration
	host := "localhost"
	port := "8080"

	// Register your routes
	http.HandleFunc("/", homeHandler)

	// Start the server
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Server is running on http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

// Define your route handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}