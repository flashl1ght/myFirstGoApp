package main

import (
	"fmt"
	"net/http"

	"github.com/flashl1ght/myFirstGoApp/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
