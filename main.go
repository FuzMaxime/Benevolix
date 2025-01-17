package main

import (
	"benevolix/config"
	"benevolix/pkg/annonce"
	"benevolix/pkg/authentification"
	"benevolix/pkg/candidature"
	"benevolix/pkg/tag"
	"benevolix/pkg/user"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	_ "benevolix/docs" // Importer les fichiers Swagger générés

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Benevolix API
// @version 1.0
// @description Benevolix API avec documentation Swagger et framework Chi.
// @host localhost:8080
// @BasePath /api/v1

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)   // Journalisation des requêtes
	router.Use(middleware.Recoverer) // Récupération des panics

	// Swagger
	router.Get("/swagger/*", httpSwagger.WrapHandler) // Documentation Swagger

	//
	router.Mount("/api/v1/auth", authentification.Routes(configuration))
	router.Mount("/api/v1/tags", tag.Routes(configuration))
	router.Mount("/api/v1/annonces", annonce.Routes(configuration))
	router.Mount("/api/v1/candidatures", candidature.Routes(configuration))
	router.Mount("/api/v1/users", user.Routes(configuration))

	return router
}

// Ouvre le swagger dans le navigateur par défaut en fonction du système d'exploitation
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start() 
	default:
		log.Println("Impossible d'ouvrir automatiquement le navigateur.")
	}

	if err != nil {
		log.Printf("Erreur lors de l'ouverture du navigateur : %v\n", err)
	}
}

func main() {
	// Génération des fichiers Swagger
	log.Println("Génération automatique de la documentation Swagger...")
	cmd := exec.Command("swag", "init")
	if err := cmd.Run(); err != nil {
		log.Fatalf("Erreur lors de l'exécution de swag init : %v", err)
	}

	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Fatalf("Erreur lors de la configuration : %v", err)
	}

	// Initialisation des routes
	router := Routes(configuration)

	// Lancement du serveur
	address := "http://localhost:8080"
	log.Printf("Serveur lancé sur %s\n", address)

	// Ouverture du Swagger dans le navigateur
	go openBrowser( address + "/swagger/index.html")

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Fatalf("Erreur du serveur : %v", err)
	}
}
