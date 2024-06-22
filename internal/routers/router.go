package routers

import (
	"kathub/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controllers.UserController) *gin.Engine{
	r := gin.Default()

	v1:=r.Group("/api/v1/users")
	{
		v1.GET("", userController.GetAll)
	}
	return r
}