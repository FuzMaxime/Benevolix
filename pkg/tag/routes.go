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

		r.Get("/", tagConfig.GetTagsHandler)
		r.Get("/{id}", tagConfig.GetTagHandler)
		r.Post("/", tagConfig.AddTagHandler)
		r.Put("/{id}", tagConfig.UpdateHandler)
		r.Delete("/{id}", tagConfig.DeleteHandler)
	})

	return router

}
