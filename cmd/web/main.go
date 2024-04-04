package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/flashl1ght/myFirstGoApp/pkg/config"
	"github.com/flashl1ght/myFirstGoApp/pkg/handlers"
	"github.com/flashl1ght/myFirstGoApp/pkg/render"
)

// change before deploying to production
const csrfAuthKey = "placerholder-csrf-auth-key"
const portNumber = ":8080"

var app config.AppConfig
var sessionManager *scs.SessionManager

// main is the main application function
func main() {

	// change to true when in production
	app.InProduction = false

	// Initialize a new session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

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
