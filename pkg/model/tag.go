package model

import (
	"errors"
	"net/http"
)

type TagRequest struct {
	Name string `json:"name" binding:"required" example:"Tag name"`
}

func (a *TagRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("name must be provided")
	}
	return nil
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
