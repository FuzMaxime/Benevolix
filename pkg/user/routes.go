package user

import (
	"benevolix/config"
	"benevolix/pkg/authentification"
	"os"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	userConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", userConfig.CreateUserHandler)

	router.Route("/", func(r chi.Router) {
		r.Use(authentification.AuthMiddleware(os.Getenv("API_Key")))
		router.Get("/", userConfig.GetAllUsersHandler)
		router.Get("/{id}", userConfig.GetByIdUserHandler)

		router.Put("/{id}", userConfig.UpdateUserHandler)

		router.Delete("/{id}", userConfig.DeleteUserHandler)

	})

	return router
}
