package dbmodel

import "gorm.io/gorm"

type TagEntry struct {
	gorm.Model
	Name     string
	Annonces []AnnonceEntry
	Users    []UserEntry
}
