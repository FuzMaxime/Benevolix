package model

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	City      string `json:"city"`
	Bio       string `json:"bio"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	if a.LastName == "" {
		return errors.New("lastname must be there")
	}
	if a.FirstName == "" {
		return errors.New("firstname must be there")
	}
	if a.Email == "" {
		return errors.New("email must be there")
	}
	if a.Password == "" {
		return errors.New("password must be there")
	}
	if a.City == "" {
		return errors.New("city must be there")
	}
	if a.Bio == "" {
		return errors.New("bio must be there")
	}
	// TODO : Implement test here
	return nil
}

type UserResponse struct {
	LastName  string        `json:"last_name"`
	FirstName string        `json:"first_name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	City      string        `json:"city"`
	Bio       string        `json:"bio"`
	Tags      []TagResponse `json:"tags"`
}
