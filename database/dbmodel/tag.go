package dbmodel

import (
	"benevolix/pkg/model"

	"gorm.io/gorm"
)

type TagEntry struct {
	gorm.Model
	Name     string
	Annonces []*AnnonceEntry `gorm:"many2many:annonce_tags"`
	Users    []*UserEntry    `gorm:"many2many:user_tags"`
}

func (tag *TagEntry) ToModel() *model.TagResponse {
	return &model.TagResponse{
		ID:   tag.ID,
		Name: tag.Name,
	}
}

type TagRepository interface {
	Create(entry *TagEntry) (*TagEntry, error)
	GetAll() ([]*TagEntry, error)
	GetById(id uint) (*TagEntry, error)
	Update(entry *TagEntry) (*TagEntry, error)
	Delete(id int) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(entry *TagEntry) (*TagEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *tagRepository) GetAll() ([]*TagEntry, error) {
	var entries []*TagEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *tagRepository) GetById(id uint) (*TagEntry, error) {
	var entrie *TagEntry
	if err := r.db.First(&entrie, id).Error; err != nil {
		return nil, err
	}
	return entrie, nil
}

func (r *tagRepository) Update(entry *TagEntry) (*TagEntry, error) {
	if err := r.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *tagRepository) Delete(id int) error {
	return r.db.Delete(&TagEntry{}, id).Error
}
