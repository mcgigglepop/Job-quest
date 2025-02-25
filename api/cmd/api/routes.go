package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create mux router
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	return mux
}
