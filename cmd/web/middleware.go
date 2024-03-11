package main

import (
	"net/http"

	"github.com/gorilla/csrf"
)

// CSRFMiddleware adds CSRF protection
func CSRFMiddleware(next http.Handler) http.Handler {
	csrfHandler := csrf.Protect(
		app.CSRFAuthKey,
		csrf.HttpOnly(true),
		csrf.Path("/"),
		csrf.SameSite(csrf.SameSiteLaxMode),
		csrf.Secure(app.InProduction),
	)

	return csrfHandler(next)
}

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}
