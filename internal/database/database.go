package database

import (
	"errors"
	"fmt"
	"kathub/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "aws-0-ap-southeast-1.pooler.supabase.com"
	port     = 6543
	user     = "postgres.urevhrzshuvxvvkuduco"
	password = "kathub12102003"
	dbName   = "postgres"
)

func DatabaseConnection() (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := AutoMigration(db); err != nil {
		return nil, fmt.Errorf("Error running migration")
	}
	return db, nil
}

func AutoMigration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		errorMessage := fmt.Sprintf("Error migrating database %v", err)
		
		return errors.New(errorMessage)
	}
	return nil
}
