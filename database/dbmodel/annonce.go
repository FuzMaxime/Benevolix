package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type AnnonceEntry struct {
	gorm.Model

	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Duration    int    `json:"duration"`
	Address     string    `json:"address"`
	IsRemote    bool      `json:"is_remote"`
	Tags        []TagEntry
	Candidature CandidatureEntry `gorm:"foreignkey:AnnonceId,references:ID"`
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
