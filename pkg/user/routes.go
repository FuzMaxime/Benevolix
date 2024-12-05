package user

import (
	"benevolix/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	userConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/create-user", userConfig.CreateUserHandler)
	router.Get("/all-users", userConfig.GetAllUsersHandler)
	router.Get("/one-user/{id}", userConfig.GetOneUserHandler)
	router.Put("/update-user/{id}", userConfig.UpdateUserHandler)
	router.Delete("/delete-user/{id}", userConfig.DeleteUserHandler)

	return router
}
