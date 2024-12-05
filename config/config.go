package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"benevolix/database"
)

type Config struct { // TODO : replace by valid repository
	// CatEntryRepository       dbmodel.CatEntryRepository
	// VisitEntryRepository     dbmodel.VisitEntryRepository
	// TreatmentEntryRepository dbmodel.TreatmentEntryRepository
}

func New() (*Config, error) {
	config := Config{}

	// Initialisation de la connexion à la base de données
	databaseSession, err := gorm.Open(sqlite.Open("clinique.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories
	// TODO : replace by valid repository
	// config.VisitEntryRepository = dbmodel.NewVisitEntryRepository(databaseSession)
	// config.TreatmentEntryRepository = dbmodel.NewTreatmentEntryRepository(databaseSession)
	// config.CatEntryRepository = dbmodel.NewCatEntryRepository(databaseSession)
	return &config, nil

}
