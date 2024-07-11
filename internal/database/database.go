package database

import (
	"os"
//	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func DatabaseConnection() {
	var err error
	sqlInfo := os.Getenv("CONN_STRING")
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("failed to connect to database")
	}
	// sqlDB, _ := DB.DB()
	// sqlDB.SetConnMaxIdleTime(time.Duration(10))
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Duration(3600))
}



