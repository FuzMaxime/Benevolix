package tag

import (
	"benevolix/config"
	"benevolix/database/dbmodel"
	"benevolix/pkg/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type TagConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TagConfig {
	return &TagConfig{configuration}
}

// GET

/*
Get all tags
*/
// GetTagsHandler gère la récupération de tous les tags
// @Summary Récupérer tous les tags
// @Description Permet de récupérer tous les tags
// @Tags Tag
// @Produce json
// @Success 200 {array} dbmodel.TagEntry
// @Failure 500 {object} map[string]string
// @Router /tags [get]
func (config *TagConfig) GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := config.TagRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	var res []model.TagResponse
	for _, tag := range tags {
		res = append(res, *tag.ToModel())
	}
	render.JSON(w, r, res)
}

/*
Get one tag by id
*/
// GetTagHandler gère la récupération d'un tag par son ID
// @Summary Récupérer un tag par son ID
// @Description Permet de récupérer un tag par son ID
// @Tags Tag
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} dbmodel.TagEntry
// @Failure 400 {object} map[string]string
// @Router /tags/{id} [get]
func (config *TagConfig) GetTagHandler(w http.ResponseWriter, r *http.Request) {
	tagID := chi.URLParam(r, "id")

	if tagID == "" {
		render.JSON(w, r, map[string]string{"error": "Tag ID is required"})
		return
	}

	id, err := strconv.ParseUint(tagID, 10, 32)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Tag ID"})
		return
	}

	tag, err := config.TagRepository.GetById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, tag.ToModel())
}

// POST

/*
Create tag
*/
// AddTagHandler gère la création d'un tag
// @Summary Créer un tag
// @Description Permet de créer un nouveau tag
// @Tags Tag
// @Accept json
// @Produce json
// @Param tag body model.TagRequest true "Tag request"
// @Success 200 {object} dbmodel.TagEntry
// @Failure 400 {object} map[string]string
// @Router /tags [post]
func (config *TagConfig) AddTagHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.TagRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "invalid tag information | " + err.Error()})
		return
	}
	tag := &dbmodel.TagEntry{
		Name: req.Name,
	}
	if _, err := config.TagRepository.Create(tag); err != nil {
		render.JSON(w, r, map[string]string{"error": "error while create tag | " + err.Error()})
		return
	}
	render.JSON(w, r, tag.ToModel())
}

// PUT

/*
Update tag
*/
// UpdateHandler gère la mise à jour d'un tag
// @Summary Mettre à jour un tag
// @Description Permet de mettre à jour un tag
// @Tags Tag
// @Accept json
// @Produce json
// @Param id path int true "Tag ID"
// @Param tag body model.TagRequest true "Tag request"
// @Success 200 {object} dbmodel.TagEntry
// @Failure 400 {object} map[string]string
// @Router /tags/{id} [put]
func (config *TagConfig) UpdateHandler(w http.ResponseWriter, r *http.Request) { // TO do
	tagId := chi.URLParam(r, "id")

	if tagId == "" {
		render.JSON(w, r, map[string]string{"error": "Tag ID is required"})
		return
	}

	id, err := strconv.ParseUint(tagId, 10, 32)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid Tag ID"})
		return
	}

	req := &model.TagRequest{}

	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	tag, err := config.TagRepository.GetById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Tag not found"})
		return
	}

	tag.Name = req.Name

	if _, err := config.TagRepository.Update(tag); err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to update Tag"})
		return
	}

	render.JSON(w, r, tag.ToModel())
}

// DELETE

// DeleteHandler gère la suppression d'un tag
// @Summary Supprimer un tag
// @Description Permet de supprimer un tag
// @Tags Tag
// @Produce json
// @Param id path int true "Tag ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /tags/{id} [delete]
func (config *TagConfig) DeleteHandler(w http.ResponseWriter, r *http.Request) { // TO do
	tagId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to get id from url path"})
		return
	}
	if err = config.TagRepository.Delete(tagId); err != nil {
		println(err.Error())
		render.JSON(w, r, map[string]string{"error": "Failed to Delete tag"})
		return
	}
	render.JSON(w, r, map[string]string{"status": "Successfully delete"})
}
