package annonce

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type AnnonceConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AnnonceConfig {
	return &AnnonceConfig{configuration}
}

// CreateAnnonceHandler gère la création d'une annonce
// @Summary Créer une annonce
// @Description Permet de créer une nouvelle annonce
// @Tags Annonce
// @Accept json
// @Produce json
// @Param annonce body model.AnnonceRequest true "Annonce request"
// @Success 200 {object} model.AnnonceResponse
// @Failure 400 {object} map[string]string
// @Router /annonce [post]
/*
{
	"title": "Titre de l'annonce",
	"description": "Description de l'annonce",
	"date": "02/01/2025",
	"duration": 2,
	"address": "Rue de la Paix 1, 1000 Lausanne",
	"is_remote": true,
	"tags": [1, 2],
	"candidature": []
}
*/
func (config *AnnonceConfig) CreateAnnonceHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.AnnonceRequest{}
	if err := render.Bind(r, req); err != nil {
		print(err.Error())
		render.JSON(w, r, map[string]string{"error": "Invalid Annonce creation request loaded"})
		return
	}

	// Check if tags exist in the database before creating the annonce
	var tags []dbmodel.TagEntry
	for _, tagId := range req.Tags {
		tag, err := config.TagRepository.GetById(tagId)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Tag not found"})
			return
		}
		tags = append(tags, *tag)
	}

	AnnonceEntry := &dbmodel.AnnonceEntry{Title: req.Title, Description: req.Description, Date: req.Date, Duration: req.Duration, Address: req.Address, IsRemote: req.IsRemote, Tags: tags}
	config.AnnonceEntryRepository.Create(AnnonceEntry)

	render.JSON(w, r, AnnonceEntry.ToModel())
}

// GetAllAnnoncesHandler gère la récupération de toutes les annonces
// @Summary Récupérer toutes les annonces
// @Description Permet de récupérer toutes les annonces
// @Tags Annonce
// @Produce json
// @Success 200 {array} model.AnnonceResponse
// @Failure 500 {object} map[string]string
// @Router /annonces [get]
func (config *AnnonceConfig) GetAllAnnoncesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.AnnonceEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	var annoncesResponse []model.AnnonceResponse
	for _, annonce := range entries {
		annoncesResponse = append(annoncesResponse, *annonce.ToModel())
	}

	render.JSON(w, r, annoncesResponse)
}

// GetOneAnnonceHandler gère la récupération d'une annonce par son ID
// @Summary Récupérer une annonce par son ID
// @Description Permet de récupérer une annonce par son ID
// @Tags Annonce
// @Produce json
// @Param id path int true "Annonce ID"
// @Success 200 {object} model.AnnonceResponse
// @Failure 400 {object} map[string]string
// @Router /annonces/{id} [get]
func (config *AnnonceConfig) GetOneAnnonceHandler(w http.ResponseWriter, r *http.Request) {
	AnnonceId := chi.URLParam(r, "id")

	entries, err := config.AnnonceEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intAnnonceId, _ := strconv.Atoi(AnnonceId)
	var AnnonceTarget *dbmodel.AnnonceEntry

	for _, Annonce := range entries {
		if Annonce.ID == uint(intAnnonceId) {
			AnnonceTarget = Annonce
		}
	}

	if AnnonceTarget == nil {
		render.JSON(w, r, map[string]string{"error": "Annonce not found"})
		return
	}
	render.JSON(w, r, AnnonceTarget.ToModel())
}

// UpdateAnnonceHandler gère la mise à jour d'une annonce
// @Summary Mettre à jour une annonce
// @Description Permet de mettre à jour une annonce
// @Tags Annonce
// @Accept json
// @Produce json
// @Param id path int true "Annonce ID"
// @Param annonce body model.AnnonceRequest true "Annonce request"
// @Success 200 {object} model.AnnonceResponse
// @Failure 400 {object} map[string]string
// @Router /annonces/{id} [put]
func (config *AnnonceConfig) UpdateAnnonceHandler(w http.ResponseWriter, r *http.Request) {
	AnnonceId := chi.URLParam(r, "id")
	intAnnonceId, err := strconv.Atoi(AnnonceId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Annonce ID"})
		return
	}

	AnnonceEntry, err := config.AnnonceEntryRepository.GetById(uint(intAnnonceId))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Annonce not found"})
		return
	}

	req := &model.AnnonceRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Annonce update request loaded"})
		return
	}
	var tags []dbmodel.TagEntry
	for _, tagId := range req.Tags {
		tag, err := config.TagRepository.GetById(tagId)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Tag not found"})
			return
		}
		tags = append(tags, *tag)
	}

	AnnonceEntry.Title = req.Title
	AnnonceEntry.Description = req.Description
	AnnonceEntry.Date = req.Date
	AnnonceEntry.Duration = req.Duration
	AnnonceEntry.Address = req.Address
	AnnonceEntry.IsRemote = req.IsRemote
	AnnonceEntry.Tags = tags

	updatedAnnonce, err := config.AnnonceEntryRepository.Update(AnnonceEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update the annonce !"})
		return
	}

	render.JSON(w, r, updatedAnnonce.ToModel())
}

// DeleteAnnonceHandler gère la suppression d'une annonce
// @Summary Supprimer une annonce
// @Description Permet de supprimer une annonce
// @Tags Annonce
// @Produce json
// @Param id path int true "Annonce ID"
// @Success 200 {string} string
// @Failure 400 {object} map[string]string
// @Router /annonces/{id} [delete]
func (config *AnnonceConfig) DeleteAnnonceHandler(w http.ResponseWriter, r *http.Request) {
	AnnonceId := chi.URLParam(r, "id")

	entries, err := config.AnnonceEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}

	intAnnonceId, err := strconv.Atoi(AnnonceId)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Annonce ID conversion"})
		return
	}
	for _, Annonce := range entries {
		if Annonce.ID == uint(intAnnonceId) {
			config.AnnonceEntryRepository.Delete(intAnnonceId)
			render.JSON(w, r, "You delete the Annonce!")
			return
		}
	}

	render.JSON(w, r, "Annonce not found")
}
