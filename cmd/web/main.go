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

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Printf("Starting application on port %s \n", portNumber)
}
