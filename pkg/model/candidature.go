package model

import (
	"errors"
	"net/http"
	"time"
)

type CandidatureRequest struct {
	UserID    uint   `json:"user_id"`
	AnnonceID uint   `json:"annonce_id"`
	Date      string `json:"date"`
	Status    string `json:"status"`
}

func (a *CandidatureRequest) Bind(r *http.Request) error {

	if a.Status != "Waiting" && a.Status != "Refused" && a.Status != "Accepted" {
		return errors.New("the status must be one of the following: Waiting or Refused or Accepted")
	}
	return nil
}

type CandidatureResponse struct {
	ID      uint      `json:"id"`
	User    uint      `json:"user"`
	Annonce uint      `json:"annonce"`
	Date    time.Time `json:"date"`
	Status  string    `json:"status"`
}
