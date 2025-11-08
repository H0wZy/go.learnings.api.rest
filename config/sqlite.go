package config

import (
	"os"

	"github.com/H0wZy/go.learnings.restapi/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db"
	dbFullPath := dbPath + "/main.db"
	// Check if db exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("db file does not exist, creating new one...")
		err = os.MkdirAll(dbPath, os.ModePerm)
		if err != nil {
			logger.Errorf("db directory creation error: %v", err)
			return nil, err
		}
		file, err := os.Create(dbFullPath)
		if err != nil {
			logger.Errorf("db file creation error: %v", err)
			return nil, err
		}
		errClose := file.Close()
		if errClose != nil {
			logger.Errorf("db file closing error: %v", errClose)
			return nil, errClose
		}
		logger.Infof("db file created")
	}
	// Create db and connect
	db, err := gorm.Open(sqlite.Open(dbFullPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite auto-migration error: %v", err)
		return nil, err
	}
	// Return db
	return db, nil
}
