package main

import (
	"net/http"

	"github.com/flashl1ght/myFirstGoApp/pkg/config"
	"github.com/gorilla/csrf"
)

var app *config.AppConfig

// NewMiddleware sets the config for middleware
func NewMiddleware(a *config.AppConfig) {
	app = a
}

// CSRFMiddleware creates middleware responsible for CSRF tokens
func CSRFMiddleware(next http.Handler) http.Handler {
	csrfHandler := csrf.Protect(
		app.CSRFAuthKey,
		csrf.HttpOnly(true),
		csrf.Path("/"),
		csrf.SameSite(csrf.SameSiteLaxMode),
		csrf.Secure(false), // remove before going to production
	)

	return csrfHandler(next)
}
