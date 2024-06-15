package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// CSRF protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Session Manager
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
