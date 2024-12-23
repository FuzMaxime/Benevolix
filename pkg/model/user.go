package model

import (
	"errors"
	"net/http"
	"net/mail"
	"unicode"
)

type UserRequest struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Phone     string `json:"phone"`
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
	for _, r := range a.LastName {
		if !unicode.IsLetter(r) {
			return errors.New("name must be charactere")
		}
	}

	for _, r := range a.FirstName {
		if !unicode.IsLetter(r) {
			return errors.New("firstname must be charactere")
		}
	}

	if a.Phone == "" && len(a.Phone) < 10 {
		return errors.New("phone must be there")
	}

	if a.Email == "" && len(a.Email) < 150 {
		return errors.New("email must be there")
	}
	_, err := mail.ParseAddress(a.Email)
	if err != nil {
		return errors.New("email must be at good format, like this : example@gmail.com")
	}

	if a.Password == "" && len(a.Password) < 100 {
		return errors.New("password must be there")
	}

	if a.City == "" && len(a.City) < 100 {
		return errors.New("city must be there")
	}
	for _, r := range a.City {
		if !unicode.IsLetter(r) {
			return errors.New("city must be charactere")
		}
	}

	if a.Bio == "" && len(a.Bio) < 150 {
		return errors.New("bio must be there")
	}
	// TODO : Implement test here
	return nil
}

type UserResponse struct {
	LastName  string        `json:"last_name"`
	FirstName string        `json:"first_name"`
	Phone     string        `json:"phone"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	City      string        `json:"city"`
	Bio       string        `json:"bio"`
	Tags      []TagResponse `json:"tags"`
}
