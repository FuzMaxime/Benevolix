package candidature

import (
	"benevolix/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	CandidatureConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/candidature", CandidatureConfig.CreateCandidatureHandler)
	router.Get("/candidature", CandidatureConfig.GetAllCandidaturesHandler)
	router.Get("/candidature/{id}", CandidatureConfig.GetOneCandidatureHandler)
	router.Put("/candidature/{id}", CandidatureConfig.UpdateCandidatureHandler)
	router.Delete("/candidature/{id}", CandidatureConfig.DeleteCandidatureHandler)

	return router
}
