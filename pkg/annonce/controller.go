package annonce

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AnnonceConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AnnonceConfig {
	return &AnnonceConfig{configuration}
}

func (config *AnnonceConfig) CreateAnnonceHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.AnnonceRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Annonce creation request loaded"})
		return
	}

	AnnonceEntry := &dbmodel.AnnonceEntry{Title: req.Title, Description: req.Description, Date: req.Date, Duration: req.Duration, Address: req.Address, IsRemote: req.IsRemote, /*Tags: req.Tags, Candidature: req.Candidature*/}
	config.AnnonceEntryRepository.Create(AnnonceEntry)

	res := &model.AnnonceResponse{Title: req.Title, Description: req.Description, Date: req.Date, Duration: req.Duration, Address: req.Address, IsRemote: req.IsRemote, /*Tags: req.Tags, Candidature: req.Candidature*/}
	render.JSON(w, r, res)
}

func (config *AnnonceConfig) GetAllAnnoncesHandler(w http.ResponseWriter, r *http.Request) {
	entries, err := config.AnnonceEntryRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

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
	render.JSON(w, r, AnnonceTarget)
}

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

	AnnonceEntry.Title = req.Title
	AnnonceEntry.Description = req.Description
	AnnonceEntry.Date = req.Date
	AnnonceEntry.Duration = req.Duration
	AnnonceEntry.Address = req.Address
	AnnonceEntry.IsRemote = req.IsRemote
	// AnnonceEntry.Tags = req.Tags
	// AnnonceEntry.Candidature = req.Candidature

	updatedAnnonce, err := config.AnnonceEntryRepository.Update(AnnonceEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update the annonce !"})
		return
	}

	render.JSON(w, r, updatedAnnonce)
}

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
