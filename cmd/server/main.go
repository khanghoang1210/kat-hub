package main

import (
	"fmt"
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/repository"
	"kathub/internal/routers"
	"kathub/internal/services"
)

func main() {
	//init db
	db, err := database.DatabaseConnection()
	fmt.Println(db.Name())
	if err != nil {
		panic(err)
	}

	// initial user instances
	userRepo := repository.NewUsersRepositoryImpl(db)
	userService := services.NewUsersServiceImpl(userRepo)
	userController := controllers.NewUserController(userService)

	r := routers.NewRouter(userController)
	r.Run()
}
