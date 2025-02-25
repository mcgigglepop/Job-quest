package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	// create mux router
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/health", app.Health)

	return mux
}
