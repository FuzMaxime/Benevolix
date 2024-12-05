package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"benevolix/database"
	"benevolix/database/dbmodel"
)

type Config struct { // TODO : replace by valid repository

	CandidatureEntryRepository dbmodel.CandidatureRepository
	TagRepository dbmodel.TagRepository
}

func New() (*Config, error) {
	config := Config{}

	// Initialisation de la connexion à la base de données
	databaseSession, err := gorm.Open(sqlite.Open("database/benevolix-database.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles
	database.Migrate(databaseSession)

	// Initialisation des repositories
	// TODO : replace by valid repository
	config.TagRepository = dbmodel.NewTagRepository(databaseSession)
	// config.VisitEntryRepository = dbmodel.NewVisitEntryRepository(databaseSession)
	// config.TreatmentEntryRepository = dbmodel.NewTreatmentEntryRepository(databaseSession)
	// config.CatEntryRepository = dbmodel.NewCatEntryRepository(databaseSession)
	config.CandidatureEntryRepository = dbmodel.NewCandidatureRepository(databaseSession)
	return &config, nil

}
