package handler

import (
	"github.com/vitorsoratto/gojobs/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger()
	db = config.GetDatabase()
}
