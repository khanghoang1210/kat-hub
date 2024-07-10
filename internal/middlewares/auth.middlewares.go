package middlewares

import (
	"fmt"
	"kathub/internal/database"
	"kathub/internal/models"
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
	fmt.Print("=============", jwtToken[1])
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
		db, err := database.DatabaseConnection()
		if err != nil {
			panic("init db Fail")
		}
		db.Where("user_name = ?", claims["username"]).First(&user)

		fmt.Print(user.UserName)
		if user.UserName == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach the request
		ctx.Set("currentUser", user)

		//Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
