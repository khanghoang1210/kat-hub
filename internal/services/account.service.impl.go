package services

import (
	"kathub/internal/repository"
	"kathub/pkg/requests"
	"kathub/pkg/responses"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AccountServiceImpl struct {
	ar repository.AccountRepository
}

func NewAccountServiceImpl(accountRepo repository.AccountRepository) AccountService {
	return &AccountServiceImpl{
		ar: accountRepo,
	}
}
func (as AccountServiceImpl) Login(req *requests.LoginReq) responses.ResponseData {
	user, err := as.ar.Login(req)

	if err != nil {
		if err.Error() == responses.StatusUserNotFound {
			return responses.ResponseData{
				StatusCode: 404,
				Message:    err.Error(),
				Data:       nil,
			}
		}	
		return responses.ResponseData{
			StatusCode: 500,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	// Compare hashed password
	success := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if success != nil {
		return responses.ResponseData{
			StatusCode: 401,
			Message:    responses.StatusAuthorizeFail,
			Data:       nil,
		}
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"sub":      user.Email,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	accessToken, errToken := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if errToken != nil {
		return responses.ResponseData{
			StatusCode: 401,
			Message:    "Failed to create token",
			Data:       nil,
		}
	}
	return responses.ResponseData{
		StatusCode: 200,
		Message:    responses.StatusSuccess,
		Data:       responses.LoginRes{
			AccessToken: accessToken,
			RefreshToken: "not yet",
		},
	}
}
