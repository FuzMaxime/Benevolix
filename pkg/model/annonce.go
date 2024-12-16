package model

import (
	"errors"
	"net/http"
	"time"
	"unicode"
)

type AnnonceRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Duration    int    `json:"duration"`
	Address     string    `json:"address"`
	IsRemote    bool      `json:"is_remote"`
	Tags        []uint    `json:"tags"`
}

func (a *AnnonceRequest) Bind(r *http.Request) error {
	if a.Title == "" && len(a.Title) < 50 {
		return errors.New("title must be there")
	}
	for _, r := range a.Title {
		if !unicode.IsLetter(r) {
			return errors.New("title must be charactere")
		}
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

	if a.Duration == "" {
		return errors.New("title must be there")
	}

	if len(a.Tags) > 0 {
		return errors.New("tag must be there")
	}
	return nil
}

type AnnonceResponse struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Date          time.Time `json:"date"`
	Duration      int    `json:"duration"`
	Address       string    `json:"address"`
	CandidatureId uint      `json:"candidature_id"`
	IsRemote      bool      `json:"is_remote"`

	Tags []TagResponse `json:"tags"`
}
