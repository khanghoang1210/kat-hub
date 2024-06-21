package server

import (
	"fmt"
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/pkg/repository"
	"kathub/pkg/services"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	//db database.DatabaseConnection
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		//db: database.New(),
	}
	db,_ := database.DatabaseConnection()
	fmt.Println(db.Name())

	userRepository:=repository.NewUsersRepositoryImpl(db)

	userService := services.NewUsersServiceImpl(userRepository)

	userController := controllers.NewTagsController(userService)
	routes := NewRoute(userController)
	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      routes,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
