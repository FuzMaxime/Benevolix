package dbmodel

import (
	"benevolix/pkg/model"
	"time"

	"gorm.io/gorm"
)

// CandidatureEntry représente une candidature dans la base de données
type CandidatureEntry struct {
	gorm.Model `swaggerignore:"true"` // Ignore gorm.Model pour Swagger
	UserID     uint                   `json:"user_id" gorm:"not null"`
	User       UserEntry              `json:"user"`
	AnnonceID  uint                   `json:"annonce_id" gorm:"not null"`
	Date       time.Time              `json:"date" gorm:"not null"`
	Status     string                 `json:"status" gorm:"not null"`

	UniqueConstraint string `gorm:"uniqueIndex:idx_user_annonce"`
}

func (candidature *CandidatureEntry) ToModel() *model.CandidatureResponse {
	return &model.CandidatureResponse{
		ID:      candidature.ID,
		UserID:  candidature.UserID,
		User:    *candidature.User.ToModel(),
		Annonce: candidature.AnnonceID,
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
	HasAlreadyApply(annonce_id, user_id int) bool
}

type candidatureRepository struct {
	db *gorm.DB
}

func NewCandidatureRepository(db *gorm.DB) CandidatureRepository {
	return &candidatureRepository{db: db}
}

func (r *candidatureRepository) Create(entry *CandidatureEntry) (*CandidatureEntry, error) {
	if entry.UserID != 0 {
		var user UserEntry
		if err := r.db.First(&user, entry.UserID).Error; err != nil {
			return nil, err
		}
		entry.User = user
	}

	if err := r.db.Create(entry).Error; err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *candidatureRepository) GetAll() ([]*CandidatureEntry, error) {
	var entries []*CandidatureEntry
	if err := r.db.Preload("User").Find(&entries).Error; err != nil {
		return nil, err
	}
	return entries, nil
}

func (r *candidatureRepository) GetById(id uint) (*CandidatureEntry, error) {
	var entry *CandidatureEntry
	if err := r.db.Preload("User").First(&entry, id).Error; err != nil {
		return nil, err
	}
	return entry, nil
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

func (r *candidatureRepository) HasAlreadyApply(annonce_id, userId int) bool {
	var candidature CandidatureEntry
	err := r.db.Where("user_id = ? AND annonce_id = ?", userId, annonce_id).First(&candidature).Error
	return err == nil
}
