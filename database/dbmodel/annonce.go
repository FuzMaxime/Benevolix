package dbmodel

import (
	"benevolix/pkg/model"
	"time"

	"gorm.io/gorm"
)

type AnnonceEntry struct {
	gorm.Model `swaggerignore:"true"` // Ignore gorm.Model pour Swagger

	Title       string             `json:"title" binding:"required" example:"Titre de l'annonce"`
	Description string             `json:"description" example:"Description de l'annonce"`
	Date        time.Time          `json:"date" binding:"required" example:"02/01/2025"`
	Duration    int                `json:"duration" binding:"required" example:"2"`
	Address     string             `json:"address" example:"Rue de la Paix 1, 1000 Lausanne"`
	IsRemote    bool               `json:"is_remote" example:"true"`
	Tags        []TagEntry         `gorm:"many2many:annonce_tags;" json:"tags" binding:"required"`
	TagIDs      []uint             `gorm:"-" json:"tags_id"`
	Candidature []CandidatureEntry `gorm:"foreignKey:AnnonceID;constraint:OnDelete:CASCADE;"`
}

func (annonce *AnnonceEntry) ToModel() *model.AnnonceResponse {
	var tags []model.TagResponse
	var candidatureResponses []model.CandidatureResponse

	for _, tag := range annonce.Tags {
		tags = append(tags, *tag.ToModel())
	}

	for _, candidature := range annonce.Candidature {
		candidatureResponses = append(candidatureResponses, model.CandidatureResponse{
			ID:      candidature.ID,
			User:    candidature.UserID,
			Annonce: annonce.ID,
			Date:    candidature.Date,
			Status:  candidature.Status,
		})
	}

	return &model.AnnonceResponse{
		ID:           annonce.ID,
		Title:        annonce.Title,
		Description:  annonce.Description,
		Date:         annonce.Date,
		Duration:     annonce.Duration,
		Address:      annonce.Address,
		Candidatures: candidatureResponses,
		IsRemote:     annonce.IsRemote,
		Tags:         tags,
	}
}

type AnnonceRepository interface {
	Create(entry *AnnonceEntry) (*AnnonceEntry, error)
	GetAll() ([]*AnnonceEntry, error)
	GetById(id uint) (*AnnonceEntry, error)
	Update(entry *AnnonceEntry) (*AnnonceEntry, error)
	Delete(id int) error
}

type annonceRepository struct {
	db *gorm.DB
}

func NewAnnonceRepository(db *gorm.DB) AnnonceRepository {
	return &annonceRepository{db: db}
}

func (r *annonceRepository) Create(entry *AnnonceEntry) (*AnnonceEntry, error) {
	if len(entry.TagIDs) > 0 {
		var tags []TagEntry
		if err := r.db.Where("id IN ?", entry.TagIDs).Find(&tags).Error; err != nil {
			return nil, err
		}
		entry.Tags = tags
	}

	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}

	return entry, nil
}

func (r *annonceRepository) GetAll() ([]*AnnonceEntry, error) {
	var entries []*AnnonceEntry
	if err := r.db.Preload("Tags").Preload("Candidature").Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *annonceRepository) GetById(id uint) (*AnnonceEntry, error) {
	var entrie *AnnonceEntry
	if err := r.db.First(&entrie, id).Error; err != nil {
		return nil, err
	}
	return entrie, nil
}

func (r *annonceRepository) Update(entry *AnnonceEntry) (*AnnonceEntry, error) {
	if err := r.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *annonceRepository) Delete(id int) error {
	return r.db.Delete(&AnnonceEntry{}, id).Error
}
