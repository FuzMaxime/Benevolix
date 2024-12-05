package dbmodel

import "gorm.io/gorm"

type UserEntry struct {
	gorm.Model
	Name        string `json:"name"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	City        string `json:"city"`
	Bio         string `json:"bio"`
	Tags        []TagEntry
	Candidature CandidatureEntry `gorm:"foreignkey:UserId,references:ID"`
}
