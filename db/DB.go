// db/db.go

package db

import (
	"BlogPost/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init initializes the database connection and sets up the schema.
func Init(connectionString string) (*gorm.DB, error) {
	// Initialize the database connection
	var err error
	db, err = gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Set up the schema by auto-migrating models
	if err := AutoMigrate(); err != nil {
		return nil, err
	}

	return db, nil
}

// AutoMigrate performs auto-migration for all models.
func AutoMigrate() error {
	if err := db.AutoMigrate(&model.BlogPost{}).Error; err != nil {
		return err
	}

	// Add more models to auto-migrate here if needed
	// e.g., db.AutoMigrate(&AnotherModel{})

	return nil
}

// GetDB returns the initialized database instance.
func GetDB() *gorm.DB {
	return db
}

// CloseDB closes the database connection.
func CloseDB() {
	if db != nil {
		db.Close()
	}
}
