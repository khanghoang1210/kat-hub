package utils

import (
	"errors"
	"kathub/pkg/responses"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) (*responses.UserResponse, error) {

	user, ok := c.Get("currentUser")
	if !ok {
		
		return nil, errors.New("Unauthorized")
	}

	currentUser, ok := user.(responses.UserResponse)
	if !ok {
		return nil, errors.New("User type assertion failed")
	}
	return &currentUser, nil

}