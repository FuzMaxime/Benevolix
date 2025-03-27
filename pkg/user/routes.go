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
		r.Get("/", userConfig.GetAllUsersHandler)
		r.Get("/{id}", userConfig.GetByIdUserHandler)

		r.Put("/{id}", userConfig.UpdateUserHandler)
		r.Put("/{id}/password", userConfig.UpdatePasswordHandler)

		r.Delete("/{id}", userConfig.DeleteUserHandler)

	})

	return router
}
