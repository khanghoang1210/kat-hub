package routers

import (
	"kathub/internal/controllers"
	"kathub/internal/middlewares"

	ws "kathub/internal/websocket"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(userController *controllers.UserController, ac *controllers.AccountController, pc *controllers.PostController) *gin.Engine {
	
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/ws", ws.NewWs)
	baseRouter := r.Group("/api/v1")
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", middlewares.AuthenMiddleware, userController.GetAll)
	userRouter.GET("/:id", userController.GetById)
	userRouter.POST("", userController.Create)
	userRouter.PUT("", userController.Update)
	userRouter.POST("/uploadAvatar",middlewares.AuthenMiddleware, userController.UploadAvatar)

	accountRouter := baseRouter.Group("/accounts")
	accountRouter.POST("/login", ac.Login)

	postRouter := baseRouter.Group("/posts")
	postRouter.POST("", middlewares.AuthenMiddleware, pc.Create)
	postRouter.GET("", pc.GetAll)
	postRouter.GET("/:id", pc.GetById)
	postRouter.PUT("/:id", middlewares.AuthenMiddleware, pc.Update)
	postRouter.DELETE("/:id", middlewares.AuthenMiddleware, pc.Delete)
	postRouter.POST("/uploadPostImage/:id", middlewares.AuthenMiddleware, pc.UploadPostImage)
	postRouter.POST("/like/:id", middlewares.AuthenMiddleware, pc.Like)
	postRouter.DELETE("/unlike/:id", middlewares.AuthenMiddleware, pc.UnLike)
	postRouter.POST("/comment/", middlewares.AuthenMiddleware, pc.CreateComment)
	postRouter.GET("/comment/:id", pc.GetCommentsByPostID)
	postRouter.PUT("/comment",middlewares.AuthenMiddleware, pc.EditComment)
	postRouter.DELETE("/comment/:id",middlewares.AuthenMiddleware, pc.DeleteComment)

	return r
}
