package main

import (
	"github.com/anucha-tk/go_basic_webapp/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(NoSurf)
	r.Use(SessionLoad)
	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/time", handlers.Repo.Time)
	r.Get("/getTime", handlers.Repo.GetTime)
	return r
}
