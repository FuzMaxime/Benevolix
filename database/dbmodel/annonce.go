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
	Duration    string    `json:"duration"`
	Address     string    `json:"address"`
	IsRemote    bool      `json:"is_remote"`
	Tags        []TagEntry
	Candidature CandidatureEntry `gorm:"foreignkey:AnnonceId,references:ID"`
}
