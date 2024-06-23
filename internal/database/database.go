package database

import (
	"errors"
	"fmt"
	"kathub/internal/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)


var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbName   = os.Getenv("DB_DATABASE")
)

func DatabaseConnection() (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

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
