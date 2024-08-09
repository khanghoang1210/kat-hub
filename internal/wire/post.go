//go:build wireinject
package wire

import (
	"kathub/internal/controllers"
	"kathub/internal/repository"
	"kathub/internal/initialize"
	"kathub/internal/services"
	"kathub/internal/database"

	"github.com/google/wire"
)

func InitPostRouterHandler() *controllers.PostController{
	wire.Build(
		database.DatabaseConnection,
		repository.NewPostRepositoryImpl, 
		initialize.CreateSession, 
		services.NewStorageServiceImpl, 
		services.NewPostServiceImpl, 
		controllers.NewPostController,
	)

	return &controllers.PostController{}
}
