package model

import (
	"errors"
	"net/http"
	"time"
)

type CandidatureRequest struct {
	UserID    uint   `json:"user_id" binding:"required"`
	AnnonceID uint   `json:"annonce_id" binding:"required"`
	Date      string `json:"date" binding:"required" example:"02/01/2025"`
	Status    string `json:"status" binding:"required" example:"Waiting"`
}

func (a *CandidatureRequest) Bind(r *http.Request) error {

	if a.Status != "Waiting" && a.Status != "Refused" && a.Status != "Accepted" {
		return errors.New("status must be one of this Waiting Refused Accepted")
	}
	// TODO : Implement test here
	return nil
}

type CandidatureResponse struct {
	ID      uint      `json:"id"`
	User    uint      `json:"user"`
	Annonce uint      `json:"annonce"`
	Date    time.Time `json:"date"`
	Status  string    `json:"status"`
}
