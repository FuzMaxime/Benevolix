package main

import (
	"benevolix/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	// TODO : replace by valid routes
	// router.Mount("/api/v1/visit", visit.Routes(configuration))
	// router.Mount("/api/v1/cat", cat.Routes(configuration))
	// router.Mount("/api/v1/treatment", treatment.Routes(configuration))
	// router.Mount("/api/v1/agecalculator", agecalculator.Routes(configuration))
	// router.Mount("/api/v1/soundidentifier", soundidentifier.Routes(configuration))
	return router
}
