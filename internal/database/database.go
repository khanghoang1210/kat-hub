package database

import (
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"launchpad.net/goamz/s3"
	"launchpad.net/goamz/aws"

)

var DB *gorm.DB
var S3Client *s3.S3

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

func S3Connection(){
	
	access := os.Getenv("S3_ACCESS_KEY")
	secret := os.Getenv("S3_SECRET_KEY")
	auth := aws.Auth{
        AccessKey: access,
        SecretKey: secret,
    }
	apsoutheast := aws.APSoutheast
	S3Client = s3.New(auth, apsoutheast)
}





