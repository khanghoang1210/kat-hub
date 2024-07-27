package main

import (
	"context"
	"fmt"
	_ "kathub/docs"
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/initialize"
	"kathub/internal/repository"
	"kathub/internal/routers"
	"kathub/internal/services"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
)

// @title   Kathub API
// @version	1.0
// @description Kathub API

// @host 	localhost:8080
// @BasePath /api/v1
func main() {
	//init db
	database.DatabaseConnection()
	fmt.Println(database.DB.Name())
	
	initialize.CreateSession()
	fmt.Println(initialize.S3Client)
	output, err := initialize.S3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("user-avatar"),
		
	})
	fmt.Println(output)
	fmt.Println(err)
	s3Service := services.NewS3ServiceImpl(initialize.S3Client)


	// initial user instance
	userRepo := repository.NewUsersRepositoryImpl(database.DB)
	userService := services.NewUsersServiceImpl(userRepo, s3Service)
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
