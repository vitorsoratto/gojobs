package config

import (
	"os"

	"github.com/vitorsoratto/gojobs/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqlite() (*gorm.DB, error) {
	logger := GetLogger()
	dbPath := "./db/database.db"

	// First check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it...")
		// If the diretory does not exist, create it
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err)
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening sqlite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error auto migrating sqlite database: %v", err)
		return nil, err
	}

	return db, nil
}
