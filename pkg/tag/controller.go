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
func (config *TagConfig) GetTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags, err := config.TagRepository.GetAll()
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, tags)
}

/*
Get one tag by id
*/
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

	entries, err := config.TagRepository.GetById(uint(id))
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to retrieve history"})
		return
	}
	render.JSON(w, r, entries)
}

// POST

/*
Create tag
*/
func (config *TagConfig) AddTagHandler(w http.ResponseWriter, r *http.Request) {
	req := &model.TagRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "invalid tag information | " + err.Error()})
		return
	}
	entries := &dbmodel.TagEntry{
		Name: req.Name,
	}
	if _, err := config.TagRepository.Create(entries); err != nil {
		render.JSON(w, r, map[string]string{"error": "error while create tag | " + err.Error()})
		return
	}
	render.JSON(w, r, entries)
}

// PUT

/*
Update tag
*/
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

	render.JSON(w, r, map[string]string{"message": "Tag updated successfully"})
}

// DELETE

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
