package tag

import (
	"benevolix/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	tagConfig := New(configuration)
	router := chi.NewRouter()

	//routes
	router.Get("/", tagConfig.GetTagsHandler)
	router.Get("/{id}", tagConfig.GetTagHandler)

	router.Post("/", tagConfig.AddTagHandler)

	router.Put("/{id}", tagConfig.UpdateHandler)

	router.Delete("/{id}", tagConfig.DeleteHandler)

	return router

}
