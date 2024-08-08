//go:build wireinject
package wire

import (
	"kathub/internal/controllers"
	"kathub/internal/repository"
	"kathub/internal/services"
	"kathub/internal/database"

	"github.com/google/wire"
)

func InitPostRouterHandler() *controllers.PostController{
	wire.Build(database.DatabaseConnection,repository.NewPostRepositoryImpl, services.NewPostServiceImpl, controllers.NewPostController)

	return &controllers.PostController{}
}
