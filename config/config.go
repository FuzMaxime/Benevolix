package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"benevolix/database"
	"benevolix/database/dbmodel"
)

type Config struct {
	UserEntryRepository dbmodel.UserRepository
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
	config.UserEntryRepository = dbmodel.NewUserRepository(databaseSession)
	return &config, nil
}
