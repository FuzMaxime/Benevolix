package authentification

import (
	"benevolix/config"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func Routes(configuration *config.Config) *chi.Mux {
	authConfig := New(configuration)
	router := chi.NewRouter()
	router.Use(AuthMiddleware("your_secret_key"))
	router.Post("/login", authConfig.Login)
	router.Get("/protected", func(w http.ResponseWriter, r *http.Request) {
		user := GetUserFromContext(r.Context())
		if user == "" {
			w.Write([]byte("Your are not register !!"))
			return
		}
		w.Write([]byte(fmt.Sprintf("Welcome, %s!", user)))
	})
	return router
}
