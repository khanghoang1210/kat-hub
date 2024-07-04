package routers

import (
	"kathub/internal/controllers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewRouter(userController *controllers.UserController, ac *controllers.AccountController, pc *controllers.PostController) *gin.Engine{
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	baseRouter:=r.Group("/api/v1")
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", userController.GetAll)
	userRouter.POST("", userController.Create)
	userRouter.PUT("", userController.Update)

	accountRouter := baseRouter.Group("/accounts")
	accountRouter.POST("/login", ac.Login)

	postRouter := baseRouter.Group("/posts")
	postRouter.POST("", pc.Create)
	return r
}