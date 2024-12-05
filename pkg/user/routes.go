package user

import (
	"benevolix/config"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	userConfig := New(configuration)
	router := chi.NewRouter()

	router.Get("/", userConfig.GetAllUsersHandler)
	router.Get("/{id}", userConfig.GetByIdUserHandler)

	router.Post("/", userConfig.CreateUserHandler)

	router.Put("/{id}", userConfig.UpdateUserHandler)

	router.Delete("/{id}", userConfig.DeleteUserHandler)

	return router
}
