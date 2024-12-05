package user

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserConfig struct {
	*config.Config
}

func New(configuration *config.Config) *UserConfig {
	return &UserConfig{configuration}
}

func (config *UserConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user creation request loaded"})
		return
	}

	userEntry := &dbmodel.UserEntry{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
		City:      req.City,
		Bio:       req.Bio,
	}

	config.UserEntryRepository.Create(userEntry)

	res := &model.UserResponse{
		LastName:  req.LastName,
		FirstName: req.FirstName,
		Email:     req.Email,
		Password:  req.Password,
		City:      req.City,
		Bio:       req.Bio,
	}
	render.JSON(w, r, res)
}

func (config *UserConfig) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.UserEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *UserConfig) GetByIdUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	intUserId, _ := strconv.Atoi(userId)

	entry, err := config.UserEntryRepository.GetById(uint(intUserId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to get user by id"})
		return
	}

	render.JSON(w, r, entry)
}

func (config *UserConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	userEntry, err := config.UserEntryRepository.GetById(uint(intUserId))
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
	userEntry.Password = req.Password
	userEntry.City = req.City
	userEntry.Bio = req.Bio

	updatedUser, err := config.UserEntryRepository.Update(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	render.JSON(w, r, updatedUser)
}

func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	intUserId, _ := strconv.Atoi(userId)

	userEntry, err := config.UserEntryRepository.GetById(uint(intUserId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	err = config.UserEntryRepository.Delete(int(userEntry.ID))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to delete user"})
		return
	}

	render.JSON(w, r, map[string]string{"message": "User deleted!"})
}
