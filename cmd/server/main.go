package main

import (
	"fmt"
	_ "kathub/docs"
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/repository"
	"kathub/internal/routers"
	"kathub/internal/services"
)

// @title   Kathub API
// @version	1.0
// @description Kathub API

// @host 	localhost:8080
// @BasePath /api/v1
func main() {
	//init db
	database.DatabaseConnection()
	fmt.Print(database.DB.Name())

	// initial user instance
	userRepo := repository.NewUsersRepositoryImpl(database.DB)
	userService := services.NewUsersServiceImpl(userRepo)
	userController := controllers.NewUserController(userService)

	// initial account instance
	ar := repository.NewAccountRepositoryImpl(database.DB)
	as := services.NewAccountServiceImpl(ar)
	ac := controllers.NewAccountController(as)
	// initial account instance
	pr := repository.NewPostRepositoryImpl(database.DB)
	ps := services.NewPostServiceImpl(pr)
	pc := controllers.NewPostController(ps)

	r := routers.NewRouter(userController, ac, pc)
	r.Run()
}
