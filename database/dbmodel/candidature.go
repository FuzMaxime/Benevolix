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
