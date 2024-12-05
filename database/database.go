package database

import (
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
	// &dbmodel.VisitEntry{},
	// &dbmodel.CatEntry{},
	// &dbmodel.TreatmentEntry{},
	)
	log.Println("Database migrated successfully")
}
