package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
    //Midlewares
	mux.Use(middleware.Recoverer)

	mux.Get("/home", app.Home);
	mux.Get("/endpt", app.Endpt);

	return mux

}