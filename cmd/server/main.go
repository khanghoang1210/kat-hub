package main

import (
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/repository"
	"kathub/internal/routers"
	"kathub/internal/services"
)

func main() {
	db,_ := database.DatabaseConnection()

	userRepo := repository.NewUsersRepositoryImpl(db)
	userService := services.NewUsersServiceImpl(userRepo)
	userController := controllers.NewUserController(userService)

	r:=routers.NewRouter(userController)
	r.Run()
}