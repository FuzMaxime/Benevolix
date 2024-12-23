package tag

import (
	"benevolix/config"
	"benevolix/pkg/authentification"
	"os"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	tagConfig := New(configuration)
	router := chi.NewRouter()

	router.Route("/api/v1", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))

		router.Get("/", tagConfig.GetTagsHandler)
		router.Get("/{id}", tagConfig.GetTagHandler)
		router.Post("/", tagConfig.AddTagHandler)
		router.Put("/{id}", tagConfig.UpdateHandler)
		router.Delete("/{id}", tagConfig.DeleteHandler)
	})

	return router

}
