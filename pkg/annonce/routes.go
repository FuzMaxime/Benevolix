package annonce

import (
	"benevolix/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	annonceConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/annonces", annonceConfig.CreateAnnonceHandler)
	router.Get("/annonces", annonceConfig.GetAllAnnoncesHandler)
	router.Get("/annonces/{id}", annonceConfig.GetOneAnnonceHandler)
	router.Put("/annonces/{id}", annonceConfig.UpdateAnnonceHandler)
	router.Delete("/annonces/{id}", annonceConfig.DeleteAnnonceHandler)

	return router
}
