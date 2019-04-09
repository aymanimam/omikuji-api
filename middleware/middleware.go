package middleware

import (
	"net/http"
)

// Middleware interface
type Middleware interface {
	ServeNext(h http.Handler) http.Handler
}

// With Add a new middleware to the handlers chain
func With(h http.Handler, ms ...Middleware) http.Handler {
	for _, m := range ms {
		h = m.ServeNext(h)
	}
	return h
}
