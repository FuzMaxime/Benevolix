package dbmodel

import (
	"time"

	"gorm.io/gorm"
)

type CandidatureEntry struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	AnnonceID uint      `json:"annonce_id"`
	Date      time.Time `json:"date"`
	Status    Status    `json:"status"`
}

type Status int

const (
	Waiting Status = iota
	Refused
	Accepted
)

func (s Status) String() string {
	switch s {
	case Waiting:
		return "Waiting"
	case Refused:
		return "Refused"
	case Accepted:
		return "Accepted"
	default:
		return "Unknown"
	}
}

type CandidatureRepository interface {
	Create(entry *AnnonceEntry) (*AnnonceEntry, error)
	GetAll() ([]*AnnonceEntry, error)
	GetById(id uint) (*AnnonceEntry, error)
	Update(entry *AnnonceEntry) (*AnnonceEntry, error)
	Delete(id int) error
}

type candidatureRepository struct {
	db *gorm.DB
}

func NewCandidatureRepository(db *gorm.DB) CandidatureRepository {
	return &candidatureRepository{db: db}
}

func (r *candidatureRepository) Create(entry *AnnonceEntry) (*AnnonceEntry, error) {
	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *candidatureRepository) GetAll() ([]*AnnonceEntry, error) {
	var entries []*AnnonceEntry
	if err := r.db.Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *candidatureRepository) GetById(id uint) (*AnnonceEntry, error) {
	var entrie *AnnonceEntry
	if err := r.db.First(&entrie, id).Error; err != nil {
		return nil, err
	}
	return entrie, nil
}

func (r *candidatureRepository) Update(entry *AnnonceEntry) (*AnnonceEntry, error) {
	if err := r.db.Save(&entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *candidatureRepository) Delete(id int) error {
	return r.db.Delete(&AnnonceEntry{}, id).Error
}
