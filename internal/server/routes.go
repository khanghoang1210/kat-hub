package server

import (
	"kathub/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)

	//r.GET("/health", s.healthHandler)

	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World by KatHub"

	c.JSON(http.StatusOK, resp)
}
func NewRoute(userController *controllers.UserController) *gin.Engine{
	service := gin.Default()
	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	router := service.Group("/api")
	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll)

	return service
}

// func (s *Server) healthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, s.db.Health())
// }
