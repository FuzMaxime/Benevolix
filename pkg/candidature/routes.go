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

		r.Post("/", CandidatureConfig.CreateCandidatureHandler)
		r.Get("/", CandidatureConfig.GetAllCandidaturesHandler)
		r.Get("/{id}", CandidatureConfig.GetOneCandidatureHandler)
		r.Put("/{id}", CandidatureConfig.UpdateCandidatureHandler)
		r.Delete("/{id}", CandidatureConfig.DeleteCandidatureHandler)
	})

	return router
}
