package candidature

import (
	"benevolix/config"
	"benevolix/pkg/authentification"
	"os"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	CandidatureConfig := New(configuration)
	router := chi.NewRouter()
	router.Route("/api/v1", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))

		router.Post("/candidature", CandidatureConfig.CreateCandidatureHandler)
		router.Get("/candidature", CandidatureConfig.GetAllCandidaturesHandler)
		router.Get("/candidature/{id}", CandidatureConfig.GetOneCandidatureHandler)
		router.Put("/candidature/{id}", CandidatureConfig.UpdateCandidatureHandler)
		router.Delete("/candidature/{id}", CandidatureConfig.DeleteCandidatureHandler)
	})

	return router
}
