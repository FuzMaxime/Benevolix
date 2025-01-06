package user

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
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
		render.JSON(w, r, map[string]string{"error": "Invalid user creation request loaded"})
		print(err.Error())
		return
	}

	userEntry := &dbmodel.UserEntry{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
		City:      req.City,
		Bio:       req.Bio,
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

	intUserId, _ := strconv.Atoi(userId)

	entry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to get user by id"})
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
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	userEntry, err := config.UserRepository.GetById(uint(intUserId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user update request loaded"})
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
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	err = config.UserRepository.Delete(int(userEntry.ID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User deleted!"})
}
