package main

import (
	"net/http"

	"github.com/ferhatyegin/goBookings/pkg/config"
	"github.com/ferhatyegin/goBookings/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))

	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux.Get("/generals-quarters", http.HandlerFunc(handlers.Repo.Generals))

	mux.Get("/majors-suite", http.HandlerFunc(handlers.Repo.Majors))

	mux.Get("/search-availability", http.HandlerFunc(handlers.Repo.Availability))
	mux.Post("/search-availability", http.HandlerFunc(handlers.Repo.PostAvailability))

	mux.Get("/contact", http.HandlerFunc(handlers.Repo.Contact))

	mux.Get("/make-reservation", http.HandlerFunc(handlers.Repo.Reservation))
	return mux
}
