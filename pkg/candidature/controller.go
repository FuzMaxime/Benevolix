package candidature

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	_ "benevolix/docs" // Importez les fichiers de documentation générés
)

type CandidatureConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CandidatureConfig {
	return &CandidatureConfig{configuration}
}

// CreateCandidatureHandler gère la création d'une candidature
// @Summary Créer une candidature
// @Description Permet de créer une nouvelle candidature
// @Tags Candidature
// @Accept json
// @Produce json
// @Param candidature body model.CandidatureRequest true "Candidature request"
// @Success 200 {object} dbmodel.CandidatureEntry
// @Failure 400 {object} map[string]string
// @Router /candidature [post]
func (config *CandidatureConfig) CreateCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.CandidatureRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature creation request loaded"})
		return
	}

	candidatureEntry := &dbmodel.CandidatureEntry{UserID: req.UserID, AnnonceID: req.AnnonceID, Date: req.Date, Status: req.Status}
	config.CandidatureRepository.Create(candidatureEntry)

	render.JSON(w, r, candidatureEntry.ToModel())
}

// GetAllCandidaturesHandler gère la récupération de toutes les candidatures
// @Summary Récupérer toutes les candidatures
// @Description Permet de récupérer toutes les candidatures
// @Tags Candidature
// @Produce json
// @Success 200 {array} dbmodel.CandidatureEntry
// @Failure 500 {object} map[string]string
// @Router /candidatures [get]
func (config *CandidatureConfig) GetAllCandidaturesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.CandidatureRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	var res []model.CandidatureResponse
	for _, entry := range entries {
		res = append(res, *entry.ToModel())
	}
	render.JSON(w, r, res)
}

// GetOneCandidatureHandler gère la récupération d'une candidature par son ID
// @Summary Récupérer une candidature par son ID
// @Description Permet de récupérer une candidature par son ID
// @Tags Candidature
// @Produce json
// @Param id path int true "Candidature ID"
// @Success 200 {object} dbmodel.CandidatureEntry
// @Failure 400 {object} map[string]string
// @Router /candidatures/{id} [get]
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

	candidature, err := config.CandidatureRepository.GetById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, candidature.ToModel())
}

// UpdateCandidatureHandler gère la mise à jour d'une candidature
// @Summary Mettre à jour une candidature
// @Description Permet de mettre à jour une candidature
// @Tags Candidature
// @Accept json
// @Produce json
// @Param id path int true "Candidature ID"
// @Param candidature body model.CandidatureRequest true "Candidature request"
// @Success 200 {object} dbmodel.CandidatureEntry
// @Failure 400 {object} map[string]string
// @Router /candidatures/{id} [put]
func (config *CandidatureConfig) UpdateCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	candidatureId := chi.URLParam(r, "id")
	intcandidatureId, err := strconv.Atoi(candidatureId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid candidature ID"})
		return
	}

	candidatureEntry, err := config.CandidatureRepository.GetById(uint(intcandidatureId))
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

	updatedCandidature, err := config.CandidatureRepository.Update(candidatureEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update candidature"})
		return
	}

	render.JSON(w, r, updatedCandidature.ToModel())
}

// DeleteCandidatureHandler gère la suppression d'une candidature
// @Summary Supprimer une candidature
// @Description Permet de supprimer une candidature
// @Tags Candidature
// @Produce json
// @Param id path int true "Candidature ID"
// @Success 200 {string} string
// @Failure 400 {object} map[string]string
// @Router /candidatures/{id} [delete]
func (config *CandidatureConfig) DeleteCandidatureHandler(w http.ResponseWriter, r *http.Request) {
	candidatureId := chi.URLParam(r, "id")

	entries, err := config.CandidatureRepository.GetAll()
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
			config.CandidatureRepository.Delete(intcandidatureId)
			render.JSON(w, r, "Your candidature is delete")
			return
		}
	}

	render.JSON(w, r, "Candidature not found")
}
