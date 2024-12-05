package candidature

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CandidatureConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CandidatureConfig {
	return &CandidatureConfig{configuration}
}

func (config *CandidatureConfig) CreateCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CandidatureRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature creation request loaded"})
		return
	}

	candidatureEntry := &dbmodel.CandidatureEntry{UserID: req.UserID, AnnonceID: req.AnnonceID, Date: req.Date, Status: req.Status}
	config.CandidatureEntryRepository.Create(candidatureEntry)

	user := config.UserEntryRepository.GetById(uint(candidatureEntry.UserID))
	annonce := config.AnnonceEntryRepository.GetById(uint(candidatureEntry.AnnonceID))

	res := &model.CandidatureResponse{User: user, Annonce: annonce, Date: req.Date, Status: req.Status}
	render.JSON(w, r, res)
}

func (config *CandidatureConfig) GetAllCandidaturesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.CandidatureEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *CandidatureConfig) GetOneCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	candidatureID := chi.URLParam(r, "id")

	if candidatureID == "" {
		render.JSON(w, r, map[string]string{"error": "Tag ID is required"})
		return
	}

	id, err := strconv.ParseUint(candidatureID, 10, 32)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Tag ID"})
		return
	}

	entries, err := config.CandidatureEntryRepository.GetById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

func (config *CandidatureConfig) UpdateCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	candidatureId := chi.URLParam(r, "id")
	intcandidatureId, err := strconv.Atoi(candidatureId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature ID"})
		return
	}

	candidatureEntry, err := config.CandidatureEntryRepository.GetById(uint(intcandidatureId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Candidature not found"})
		return
	}

	req := &model.CandidatureRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature update request loaded"})
		return
	}

	candidatureEntry.UserID = req.UserID
	candidatureEntry.AnnonceID = req.AnnonceID
	candidatureEntry.Date = req.Date
	candidatureEntry.Status = req.Status

	updatedCandidature, err := config.CandidatureEntryRepository.Update(candidatureEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update candidature"})
		return
	}

	render.JSON(w, r, updatedCandidature)
}

func (config *CandidatureConfig) DeleteCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	candidatureId := chi.URLParam(r, "id")

	entries, err := config.CandidatureEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intcandidatureId, err := strconv.Atoi(candidatureId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature ID conversion"})
		return
	}
	for _, candidature := range entries {
		if candidature.ID == uint(intcandidatureId) {
			config.CandidatureEntryRepository.Delete(intcandidatureId)
			render.JSON(w, r, "Your candidature is delete")
			return
		}
	}

	render.JSON(w, r, "Candidature not found")
}
