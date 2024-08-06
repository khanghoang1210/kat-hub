package main

import (
	"fmt"
	"kathub/internal/database"
	"kathub/internal/models"
)

func main() {
	// db,err := database.DatabaseConnection()
	// if  err != nil {
	// 	errorMessage := fmt.Sprintf("Error migrating database %v", err)

	// 	panic(errorMessage)
	// }
	//db.Debug().AutoMigrate(&models.User{});
	 // Auto migrate the User struct to add new columns
	 database.DatabaseConnection()
	 if err := database.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}, &models.Comment{}); err != nil {
        fmt.Printf("failed to auto migrate: %v", err)
    }
}
