//go:build wireinject
package wire

import (
	"kathub/internal/controllers"
	"kathub/internal/repository"
	"kathub/internal/services"
	"kathub/internal/database"

	"github.com/google/wire"
)

func InitAccountRouterHandler() *controllers.AccountController{
	wire.Build(database.DatabaseConnection,repository.NewAccountRepositoryImpl, services.NewAccountServiceImpl, controllers.NewAccountController)

	return &controllers.AccountController{}
}
