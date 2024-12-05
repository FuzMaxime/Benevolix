package model

import (
	"errors"
	"net/http"
	"time"
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
	if a.Title == "" {
		return errors.New("title must be there")
	}
	// TODO : Implement test here
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
