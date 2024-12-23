package annonce

import (
	"benevolix/config"
	"benevolix/pkg/authentification"
	"os"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	annonceConfig := New(configuration)
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))

		router.Post("/", annonceConfig.CreateAnnonceHandler)
		router.Get("/", annonceConfig.GetAllAnnoncesHandler)
		router.Get("/{id}", annonceConfig.GetOneAnnonceHandler)
		router.Put("/{id}", annonceConfig.UpdateAnnonceHandler)
		router.Delete("/{id}", annonceConfig.DeleteAnnonceHandler)
	})

	return router
}
