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
	router.Route("/", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))

		r.Post("/candidature", CandidatureConfig.CreateCandidatureHandler)
		r.Get("/candidature", CandidatureConfig.GetAllCandidaturesHandler)
		r.Get("/candidature/{id}", CandidatureConfig.GetOneCandidatureHandler)
		r.Put("/candidature/{id}", CandidatureConfig.UpdateCandidatureHandler)
		r.Delete("/candidature/{id}", CandidatureConfig.DeleteCandidatureHandler)
	})

	return router
}
