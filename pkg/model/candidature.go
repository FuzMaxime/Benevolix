package model

import (
	"errors"
	"net/http"
	"time"
)

type CandidatureRequest struct {
	UserID    uint      `json:"user_id"`
	AnnonceID uint      `json:"annonce_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"`
}

func (a *CandidatureRequest) Bind(r *http.Request) error {

	if a.Status != "Waiting" && a.Status != "Refused" && a.Status != "Accepted" {
		return errors.New("status must be one of this Waiting Refused Accepted")
	}
	// TODO : Implement test here
	return nil
}

type CandidatureResponse struct {
	ID      uint            `json:"id"`
	User    UserResponse    `json:"user"`
	Annonce AnnonceResponse `json:"annonce"`
	Date    time.Time       `json:"date"`
	Status  string          `json:"status"`
}
