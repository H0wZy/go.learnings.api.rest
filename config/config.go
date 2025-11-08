package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDb() error {
	return errors.New("not implemented")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
