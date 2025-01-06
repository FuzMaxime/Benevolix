# Benevolix

Bénévolat entre Gaulois

## Description

Benevolix est une API pour gérer les annonces, candidatures, utilisateurs et tags pour une plateforme de bénévolat. L'API est documentée avec Swagger et utilise le framework Chi pour la gestion des routes.

## Prérequis

- Go 1.16 ou supérieur
- [Swag](https://github.com/swaggo/swag) pour générer la documentation Swagger

## Installation

1. Clonez le dépôt :

   ```sh
   git clone https://github.com/FuzMaxime/Benevolix.git
   ```

2. Installez les dépendances :

   ```sh
   go mod download
   ```

3. Installez Swag :
   ```sh
   go install github.com/swaggo/swag/cmd/swag@latest
   export PATH=$PATH:$(go env GOPATH)/bin
   ```

## Génération de la documentation Swagger

Pour générer la documentation Swagger, exécutez la commande suivante :

```sh
swag init
```

**Celle-ci est générée automatiquement au lancement du projet !**

## Lancer le projet

Premier lancer :

```sh
go mod tidy
go run main.go
```

Pour lancer le projet, exécutez la commande suivante :

```sh
go run main.go
```

L'API sera disponible à l'adresse http://localhost:8080.

## Endpoints

- `/api/v1/auth` - Routes d'authentification
- `/api/v1/tags` - Routes des tags
- `/api/v1/annonces` - Routes des annonces
- `/api/v1/candidatures` - Routes des candidatures
- `/api/v1/users` - Routes des utilisateurs

## Documentation Swagger

La documentation Swagger est disponible à l'adresse http://localhost:8080/swagger/index.html.
