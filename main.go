package main

import (
	"benevolix/config"
	"benevolix/pkg/user"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/user", user.Routes(configuration))
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
	log.Fatal(http.ListenAndServe(":8080", router))
}
