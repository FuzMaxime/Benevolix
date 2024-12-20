package dbmodel

import (
	"benevolix/pkg/model"
	"time"

	"gorm.io/gorm"
)

type CandidatureEntry struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	AnnonceID uint      `json:"annonce_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}

func (candidature *CandidatureEntry) ToModel() *model.CandidatureResponse {
	var r *candidatureRepository

	var userEntrie *UserEntry
	if err := r.db.First(&userEntrie, candidature.UserID).Error; err != nil {
	}

	var annonceEntrie *AnnonceEntry
	if err := r.db.First(&annonceEntrie, candidature.AnnonceID).Error; err != nil {
	}
	return &model.CandidatureResponse{
		User:    *userEntrie.ToModel(),
		Annonce: *annonceEntrie.ToModel(),
		Date:    candidature.Date,
		Status:  candidature.Status,
	}
}

type CandidatureRepository interface {
	Create(entry *CandidatureEntry) (*CandidatureEntry, error)
	GetAll() ([]*CandidatureEntry, error)
	GetById(id uint) (*CandidatureEntry, error)
	Update(entry *CandidatureEntry) (*CandidatureEntry, error)
	Delete(id int) error
}

type candidatureRepository struct {
	db *gorm.DB
}

func NewCandidatureRepository(db *gorm.DB) CandidatureRepository {
	return &candidatureRepository{db: db}
}

func (r *candidatureRepository) Create(entry *CandidatureEntry) (*CandidatureEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *candidatureRepository) GetAll() ([]*CandidatureEntry, error) {
	var entries []*CandidatureEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *candidatureRepository) GetById(id uint) (*CandidatureEntry, error) {
	var entrie *CandidatureEntry
	if err := r.db.First(&entrie, id).Error; err != nil {
		return nil, err
	}
	return entrie, nil
}

func (r *candidatureRepository) Update(entry *CandidatureEntry) (*CandidatureEntry, error) {
	if err := r.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *candidatureRepository) Delete(id int) error {
	return r.db.Delete(&CandidatureEntry{}, id).Error
}
