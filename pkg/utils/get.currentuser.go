package utils

import "github.com/gin-gonic/gin"

func GetUserProfile(c *gin.Context) any {

	user, _ := c.Get("currentUser")

	c.JSON(200, gin.H{
		"user": user,
	})
	return user
}