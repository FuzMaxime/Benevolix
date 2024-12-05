package dbmodel

import (
	"benevolix/pkg/model"
	"time"

	"gorm.io/gorm"
)

type AnnonceEntry struct {
	gorm.Model

	Title       string            `json:"title"`
	Description string            `json:"description"`
	Date        time.Time         `json:"date"`
	Duration    string            `json:"duration"`
	Address     string            `json:"address"`
	IsRemote    bool              `json:"is_remote"`
	Tags        []*TagEntry       `gorm:"many2many:annonce_tags"`
	Candidature *CandidatureEntry `gorm:"foreignkey:AnnonceID;references:ID"`
}

func (annonce *AnnonceEntry) ToModel() *model.AnnonceResponse {
	var tags []model.TagResponse
	for _, tag := range annonce.Tags {
		tags = append(tags, *tag.ToModel())
	}
	return &model.AnnonceResponse{
		Title:         annonce.Title,
		Description:   annonce.Description,
		Date:          annonce.Date,
		Duration:      annonce.Duration,
		Address:       annonce.Address,
		CandidatureId: annonce.Candidature.ID,
		IsRemote:      annonce.IsRemote,
		Tags:          tags,
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
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *annonceRepository) GetAll() ([]*AnnonceEntry, error) {
	var entries []*AnnonceEntry
	if err := r.db.Find(&entries).Error; err != nil {
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
