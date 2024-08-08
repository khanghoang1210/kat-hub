//go:build wireinject

package wire

import (
	"kathub/internal/controllers"
	"kathub/internal/database"
	"kathub/internal/initialize"
	"kathub/internal/repository"
	"kathub/internal/services"

	"github.com/google/wire"
)

func InitUserRouterHandler() *controllers.UserController {
	
	wire.Build(
		database.DatabaseConnection,
		repository.NewUsersRepositoryImpl,
		initialize.CreateSession,
		services.NewStorageServiceImpl,
		services.NewUsersServiceImpl,
		
		controllers.NewUserController,
	)
	return &controllers.UserController{}
}