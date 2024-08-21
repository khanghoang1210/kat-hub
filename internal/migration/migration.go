package main

import (
	"fmt"
	"kathub/internal/database"
	"kathub/internal/models"
)

func main() {
	 database.DatabaseConnection()
	 if err := database.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}, &models.Comment{}, &models.Friend{}); err != nil {
        fmt.Printf("failed to auto migrate: %v", err)
    }
}
