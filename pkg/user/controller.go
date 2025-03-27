package user

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type UserConfig struct {
	*config.Config
}

func New(configuration *config.Config) *UserConfig {
	return &UserConfig{configuration}
}

// CreateUserHandler gère la création d'un utilisateur
// @Summary Créer un utilisateur
// @Description Permet de créer un nouvel utilisateur
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.UserRequest true "User request"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Router /user [post]
func (config *UserConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	var tags []dbmodel.TagEntry
	for _, tag := range req.Tags {
		tagEntry, err := config.TagRepository.GetById(tag)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "tag not found"})
			return
		}
		tags = append(tags, *tagEntry)
	}

	userEntry := &dbmodel.UserEntry{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		City:      req.City,
		Bio:       req.Bio,
		Tags:      tags,
	}

	if _, err := config.UserRepository.Create(userEntry); err != nil {
		render.JSON(w, r, map[string]string{"error": "User already registered"})
		print(err.Error())
		return
	}
	render.JSON(w, r, userEntry.ToModel())
}

// GetAllUsersHandler gère la récupération de tous les utilisateurs
// @Summary Récupérer tous les utilisateurs
// @Description Permet de récupérer tous les utilisateurs
// @Tags User
// @Produce json
// @Success 200 {array} dbmodel.UserEntry
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (config *UserConfig) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := config.UserRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	var res []model.UserResponse
	for _, user := range users {
		res = append(res, *user.ToModel())
	}
	render.JSON(w, r, res)
}

// GetByIdUserHandler gère la récupération d'un utilisateur par son ID
// @Summary Récupérer un utilisateur par son ID
// @Description Permet de récupérer un utilisateur par son ID
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dbmodel.UserEntry
// @Failure 400 {object} map[string]string
// @Router /users/{id} [get]
func (config *UserConfig) GetByIdUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	entry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	// Vérifiez si l'utilisateur est vide ou nil
	if entry == nil {
		w.WriteHeader(http.StatusNoContent) // 204 No Content
		return
	}

	render.JSON(w, r, entry.ToModel())
}

// UpdateUserHandler gère la mise à jour d'un utilisateur
// @Summary Mettre à jour un utilisateur
// @Description Permet de mettre à jour un utilisateur
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.UserRequest true "User request"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} map[string]string
// @Router /users/{id} [put]
func (config *UserConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	userEntry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	userEntry.LastName = req.LastName
	userEntry.FirstName = req.FirstName
	userEntry.Email = req.Email
	userEntry.Phone = req.Phone
	userEntry.Password = req.Password
	userEntry.City = req.City
	userEntry.Bio = req.Bio

	updatedUser, err := config.UserRepository.Update(userEntry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}
	render.JSON(w, r, updatedUser.ToModel())
}

// DeleteUserHandler gère la suppression d'un utilisateur
// @Summary Supprimer un utilisateur
// @Description Permet de supprimer un utilisateur
// @Tags User
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {string} string
// @Failure 400 {object} map[string]string
// @Router /users/{id} [delete]
func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	intUserId, _ := strconv.Atoi(userId)

	userEntry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	err = config.UserRepository.Delete(int(userEntry.ID))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}

	w.WriteHeader(http.StatusOK)
	render.JSON(w, r, map[string]string{"message": "User deleted"})
}

// UpdatePasswordHandler gère la mise à jour du mot de passe d'un utilisateur
// @Summary Mettre à jour le mot de passe d'un utilisateur
// @Description Permet de mettre à jour le mot de passe d'un utilisateur après vérification du mot de passe actuel
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param password body map[string]string true "Password request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /users/{id}/password [put]
func (config *UserConfig) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Récupérer l'utilisateur depuis la base de données
	userEntry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	// Lire le corps de la requête
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request"})
		return
	}

	currentPassword, ok := req["currentPassword"]
	if !ok || currentPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Current password is required"})
		return
	}

	newPassword, ok := req["newPassword"]
	if !ok || newPassword == "" {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "New password is required"})
		return
	}

	// Vérifier que le mot de passe actuel correspond
	if err := bcrypt.CompareHashAndPassword([]byte(userEntry.Password), []byte(currentPassword)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		render.JSON(w, r, map[string]string{"error": "Current password is incorrect"})
		return
	}

	// Hashage du nouveau mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to hash new password"})
		return
	}

	// Mise à jour du mot de passe dans la base de données
	userEntry.Password = string(hashedPassword)
	if _, err := config.UserRepository.Update(userEntry); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Failed to update password"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "Password updated successfully"})
}
