package model

import (
	"errors"
	"net/http"
	"time"
)

type AnnonceRequest struct {
	Title       string    `json:"title" binding:"required" example:"Titre de l'annonce"`
	Description string    `json:"description" example:"Description de l'annonce" binding:"required"`
	Date        time.Time `json:"date" binding:"required" example:"02/01/2025" time_format:"02/01/2006"`
	Duration    int       `json:"duration" binding:"required" example:"2"`
	Address     string    `json:"address" binding:"required" example:"Rue de la Paix 1, 1000 Lausanne"`
	IsRemote    bool      `json:"is_remote" example:"true" binding:"required"`
	Tags        []uint    `json:"tags"`
}

func (a *AnnonceRequest) Bind(r *http.Request) error {
	if a.Title == "" && len(a.Title) < 50 {
		return errors.New("title must be there")
	}

	if a.Description == "" && len(a.Description) < 200 {
		return errors.New("description must be there")
	}

	now := time.Now()
	oneYearLater := now.AddDate(1, 0, 0)
	if a.Date.Before(now) {
		return errors.New("the date must be after now")
	} else if a.Date.After(oneYearLater) {
		return errors.New("the date must be before today + 1 years")
	}

	if a.Duration < 0 {
		return errors.New("duration must be provided")
	}

	return nil
}

type AnnonceResponse struct {
	ID          uint      `json:"id"`
	OwnerID     uint      `json:"owner_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Duration    int       `json:"duration"`
	Address     string    `json:"address"`
	IsRemote    bool      `json:"is_remote"`

	Tags         []TagResponse         `json:"tags"`
	Candidatures []CandidatureResponse `json:"candidatures"`
}
