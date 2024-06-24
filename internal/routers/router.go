package routers

import (
	"kathub/internal/controllers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewRouter(userController *controllers.UserController) *gin.Engine{
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	baseRouter:=r.Group("/api/v1")
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", userController.GetAll)
	userRouter.POST("", userController.Create)
	return r
}