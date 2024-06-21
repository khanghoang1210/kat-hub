package main

import (
	"fmt"
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/server"
	"kathub/pkg/repository"
	"kathub/pkg/services"
	"net/http"
	"time"
)

func init(){

}
func main() {

	db,_ := database.DatabaseConnection()
	fmt.Println(db.Name())

	userRepository:=repository.NewUsersRepositoryImpl(db)

	userService := services.NewUsersServiceImpl(userRepository)

	userController := controllers.NewTagsController(userService)
	routes := server.NewRoute(userController)
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	server.ListenAndServe()
}
