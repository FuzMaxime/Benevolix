package model

import (
	"errors"
	"net/http"
	"net/mail"
	"unicode"
)

type UserRequest struct {
	LastName  string `json:"last_name" binding:"required" example:"Nom de famille"`
	FirstName string `json:"first_name" binding:"required" example:"Prénom"`
	Phone     string `json:"phone" example:"0791234567" binding:"required"`
	Email     string `json:"email" example:"example@example.com" binding:"required"`
	Password  string `json:"password" example:"password" binding:"required"`
	City      string `json:"city" example:"Nantes" binding:"required"`
	Bio       string `json:"bio" example:"Je suis un étudiant en informatique" binding:"required"`
}

func (a *UserRequest) Bind(r *http.Request) error {
	for _, r := range a.LastName {
		if !unicode.IsLetter(r) {
			return errors.New("last name must be charactere")
		}
	}

	for _, r := range a.FirstName {
		if !unicode.IsLetter(r) {
			return errors.New("first name must be charactere")
		}
	}

	if a.Phone == "" {
		return errors.New("phone must be provided")
	}

	if a.Email == "" {
		return errors.New("email must be provided")
	}
	_, err := mail.ParseAddress(a.Email)
	if err != nil {
		return errors.New("email must be at good format, like this : example@gmail.com")
	}

	if a.Password == "" {
		return errors.New("password must be provided")
	}

	if a.City == "" {
		return errors.New("city must be provided")
	}
	for _, r := range a.City {
		if !unicode.IsLetter(r) {
			return errors.New("city must be charactere")
		}
	}

	if a.Bio == "" {
		return errors.New("biography must be provided")
	}
	return nil
}

type UserResponse struct {
	ID        uint          `json:"id"`
	LastName  string        `json:"last_name"`
	FirstName string        `json:"first_name"`
	Phone     string        `json:"phone"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	City      string        `json:"city"`
	Bio       string        `json:"bio"`
	Tags      []TagResponse `json:"tags"`
}
