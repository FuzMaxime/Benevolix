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
	router.Route("/", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))

		r.Post("/", annonceConfig.CreateAnnonceHandler)
		r.Get("/", annonceConfig.GetAllAnnoncesHandler)
		r.Get("/{id}", annonceConfig.GetOneAnnonceHandler)
		r.Put("/{id}", annonceConfig.UpdateAnnonceHandler)
		r.Delete("/{id}", annonceConfig.DeleteAnnonceHandler)
	})

	return router
}
