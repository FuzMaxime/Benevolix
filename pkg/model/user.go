package model

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	City      string `json:"city"`
	Bio       string `json:"bio"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("name must be there")
	}
	// TODO : Implement test here
	return nil
}

type UserResponse struct {
	Name      string        `json:"name"`
	FirstName string        `json:"first_name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	City      string        `json:"city"`
	Bio       string        `json:"bio"`
	Tags      []TagResponse `json:"tags"`
}
