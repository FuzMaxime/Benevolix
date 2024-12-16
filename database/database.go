package database

import (
	"benevolix/database/dbmodel"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&dbmodel.TagEntry{},
		&dbmodel.AnnonceEntry{},
		&dbmodel.CandidatureEntry{},
		&dbmodel.UserEntry{},
	)
	log.Println("Database migrated successfully")
}
