package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flashl1ght/myFirstGoApp/pkg/config"
	"github.com/flashl1ght/myFirstGoApp/pkg/handlers"
	"github.com/flashl1ght/myFirstGoApp/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
