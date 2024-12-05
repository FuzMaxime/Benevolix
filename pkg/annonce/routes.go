package annonce

import (
	"benevolix/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	annonceConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", annonceConfig.CreateAnnonceHandler)
	router.Get("/", annonceConfig.GetAllAnnoncesHandler)
	router.Get("/{id}", annonceConfig.GetOneAnnonceHandler)
	router.Put("/{id}", annonceConfig.UpdateAnnonceHandler)
	router.Delete("/{id}", annonceConfig.DeleteAnnonceHandler)

	return router
}
