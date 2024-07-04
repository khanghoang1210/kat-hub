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
	db, err := database.DatabaseConnection()
	fmt.Println(db.Name())
	if err != nil {
		panic(err)
	}

	// initial user instance
	userRepo := repository.NewUsersRepositoryImpl(db)
	userService := services.NewUsersServiceImpl(userRepo)
	userController := controllers.NewUserController(userService)

	// initial account instance
	ar := repository.NewAccountRepositoryImpl(db)
	as := services.NewAccountServiceImpl(ar)
	ac := controllers.NewAccountController(as)
	// initial account instance
	pr := repository.NewPostRepositoryImpl(db)
	ps := services.NewPostServiceImpl(pr)
	pc := controllers.NewPostController(ps)

	r := routers.NewRouter(userController, ac, pc)
	r.Run()
}
