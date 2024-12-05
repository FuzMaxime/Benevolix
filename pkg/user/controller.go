package user

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodel"
	"vet-clinic-api/pkg/model"

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

	userEntry := &dbmodel.UserEntry{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	config.UserEntryRepository.Create(userEntry)

	res := &model.UserResponse{Name: req.Name, Age: req.Age, Race: req.Race, Weight: req.Weight}
	render.JSON(w, r, res)
}

func (config *UserConfig) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.UserEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *UserConfig) GetOneUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	entries, err := config.UserEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intUserId, _ := strconv.Atoi(userId)
	var userTarget *dbmodel.UserEntry

	for _, user := range entries {
		if user.ID == uint(intUserId) {
			userTarget = user
		}
	}

	if userTarget == nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	render.JSON(w, r, userTarget)
}

func (config *UserConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID"})
		return
	}

	userEntry, err := config.UserEntryRepository.FindByID(uint(intUserId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}

	req := &model.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user update request loaded"})
		return
	}

	userEntry.Name = req.Name
	userEntry.Age = req.Age
	userEntry.Race = req.Race
	userEntry.Weight = req.Weight

	updatedUser, err := config.UserEntryRepository.Update(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update user"})
		return
	}

	render.JSON(w, r, updatedUser)
}

func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	entries, err := config.UserEntryRepository.FindAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intUserId, err := strconv.Atoi(userId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid user ID conversion"})
		return
	}
	for _, user := range entries {
		if user.ID == uint(intUserId) {
			config.UserEntryRepository.Delete(user)
			render.JSON(w, r, "Oups, we have kill your user!")
			return
		}
	}

	render.JSON(w, r, "User not found")
}
