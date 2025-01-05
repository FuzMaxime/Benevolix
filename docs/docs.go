// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/annonce": {
            "post": {
                "description": "Permet de créer une nouvelle annonce",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Annonce"
                ],
                "summary": "Créer une annonce",
                "parameters": [
                    {
                        "description": "Annonce request",
                        "name": "annonce",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AnnonceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AnnonceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/annonces": {
            "get": {
                "description": "Permet de récupérer toutes les annonces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Annonce"
                ],
                "summary": "Récupérer toutes les annonces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.AnnonceResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/annonces/{id}": {
            "get": {
                "description": "Permet de récupérer une annonce par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Annonce"
                ],
                "summary": "Récupérer une annonce par son ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Annonce ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AnnonceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Permet de mettre à jour une annonce",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Annonce"
                ],
                "summary": "Mettre à jour une annonce",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Annonce ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Annonce request",
                        "name": "annonce",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AnnonceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AnnonceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Permet de supprimer une annonce",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Annonce"
                ],
                "summary": "Supprimer une annonce",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Annonce ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/candidature": {
            "post": {
                "description": "Permet de créer une nouvelle candidature",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidature"
                ],
                "summary": "Créer une candidature",
                "parameters": [
                    {
                        "description": "Candidature request",
                        "name": "candidature",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CandidatureRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.CandidatureEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/candidatures": {
            "get": {
                "description": "Permet de récupérer toutes les candidatures",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidature"
                ],
                "summary": "Récupérer toutes les candidatures",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbmodel.CandidatureEntry"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/candidatures/{id}": {
            "get": {
                "description": "Permet de récupérer une candidature par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidature"
                ],
                "summary": "Récupérer une candidature par son ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Candidature ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.CandidatureEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Permet de mettre à jour une candidature",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidature"
                ],
                "summary": "Mettre à jour une candidature",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Candidature ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Candidature request",
                        "name": "candidature",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CandidatureRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.CandidatureEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Permet de supprimer une candidature",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidature"
                ],
                "summary": "Supprimer une candidature",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Candidature ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Permet à un utilisateur de se connecter avec ses identifiants.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentification"
                ],
                "summary": "Connexion de l'utilisateur",
                "parameters": [
                    {
                        "description": "Login payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/authentification.LoginPayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "description": "Permet de récupérer tous les tags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Récupérer tous les tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbmodel.TagEntry"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Permet de créer un nouveau tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Créer un tag",
                "parameters": [
                    {
                        "description": "Tag request",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.TagEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/tags/{id}": {
            "get": {
                "description": "Permet de récupérer un tag par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Récupérer un tag par son ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.TagEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Permet de mettre à jour un tag",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Mettre à jour un tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Tag request",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.TagEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Permet de supprimer un tag",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tag"
                ],
                "summary": "Supprimer un tag",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Tag ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/user": {
            "post": {
                "description": "Permet de créer un nouvel utilisateur",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Créer un utilisateur",
                "parameters": [
                    {
                        "description": "User request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Permet de récupérer tous les utilisateurs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Récupérer tous les utilisateurs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dbmodel.UserEntry"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Permet de récupérer un utilisateur par son ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Récupérer un utilisateur par son ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dbmodel.UserEntry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Permet de mettre à jour un utilisateur",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Mettre à jour un utilisateur",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Permet de supprimer un utilisateur",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Supprimer un utilisateur",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "authentification.LoginPayload": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dbmodel.AnnonceEntry": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "candidature": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbmodel.CandidatureEntry"
                    }
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "is_remote": {
                    "type": "boolean"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbmodel.TagEntry"
                    }
                },
                "tags_id": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dbmodel.CandidatureEntry": {
            "type": "object",
            "properties": {
                "annonce_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dbmodel.TagEntry": {
            "type": "object",
            "properties": {
                "annonces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbmodel.AnnonceEntry"
                    }
                },
                "name": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbmodel.UserEntry"
                    }
                }
            }
        },
        "dbmodel.UserEntry": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "candidature": {
                    "$ref": "#/definitions/dbmodel.CandidatureEntry"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dbmodel.TagEntry"
                    }
                }
            }
        },
        "model.AnnonceRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "is_remote": {
                    "type": "boolean"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.AnnonceResponse": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "candidatures": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.CandidatureResponse"
                    }
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "duration": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "is_remote": {
                    "type": "boolean"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TagResponse"
                    }
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.CandidatureRequest": {
            "type": "object",
            "properties": {
                "annonce_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "model.CandidatureResponse": {
            "type": "object",
            "properties": {
                "annonce": {
                    "$ref": "#/definitions/model.AnnonceResponse"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.UserResponse"
                }
            }
        },
        "model.TagRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.TagResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.UserRequest": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.TagResponse"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Benevolix API",
	Description:      "Benevolix API avec documentation Swagger et framework Chi.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}