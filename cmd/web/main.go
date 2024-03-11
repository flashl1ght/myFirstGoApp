package main

import (
	"log"
	"net/http"

	"github.com/flashl1ght/myFirstGoApp/pkg/config"
	"github.com/flashl1ght/myFirstGoApp/pkg/handlers"
	"github.com/flashl1ght/myFirstGoApp/pkg/render"
)

const portNumber = ":8080"

// change before deploying to production
const csrfAuthKey = "placerholder-csrf-auth-key"

// main is the main application function
func main() {
	var app config.AppConfig

	// cache templates
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache

	// read CSRF auth key
	app.CSRFAuthKey = []byte(csrfAuthKey)

	// pass AppConfig
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	NewMiddleware(&app)

	// server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Printf("Starting application on port %s \n", portNumber)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
