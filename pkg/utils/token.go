package utils

import (
	"fmt"
	"kathub/internal/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  user.UserName,
		"sub":       user.Email,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"authorize": true,
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func TokenValid(ctx *gin.Context) error {
	tokenString := ctx.GetHeader("Authorization")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	return nil
}