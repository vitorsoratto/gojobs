package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDatabase() error {
	var err error
	db, err = InitSqlite()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}

func GetLogger() *Logger {
	logger = NewLogger()
	return logger
}
