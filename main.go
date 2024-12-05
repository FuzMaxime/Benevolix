package main

import (
	"benevolix/config"
	"benevolix/pkg/tag"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()
	// TODO : replace by valid routes
	router.Mount("/api/v1/tag", tag.Routes(configuration))
	// router.Mount("/api/v1/visit", visit.Routes(configuration))
	// router.Mount("/api/v1/cat", cat.Routes(configuration))
	// router.Mount("/api/v1/treatment", treatment.Routes(configuration))
	// router.Mount("/api/v1/agecalculator", agecalculator.Routes(configuration))
	// router.Mount("/api/v1/soundidentifier", soundidentifier.Routes(configuration))
	return router
}

func main() {
	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	// Initialisation des routes
	router := Routes(configuration)

	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
