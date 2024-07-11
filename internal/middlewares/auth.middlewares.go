package middlewares

import (
	"fmt"
	"kathub/internal/database"
	"kathub/internal/models"
	"kathub/pkg/responses"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenMiddleware(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")

	jwtToken := strings.Split(tokenString, " ")

	if len(jwtToken) != 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/validate it
	bearerToken, _ := jwt.Parse(jwtToken[1], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	fmt.Print(bearerToken.Valid)

	if claims, ok := bearerToken.Claims.(jwt.MapClaims); ok && bearerToken.Valid {

		// Check the expiry date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user with token Subject
		var user models.User

		database.DB.First(&user, claims["sub"])

		if user.UserName == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		currentUser := responses.UserResponse{
			Id: user.Id,
			UserName: user.UserName,
			FullName: user.FullName,
			Email: user.Email,
			Gender: user.Gender,
			AvatarUrl: user.AvatarUrl,
			CreatedAt: user.CreatedAt,
		}
		// Attach the request
		ctx.Set("currentUser", currentUser)

		//Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
