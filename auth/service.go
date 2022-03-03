package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GenerateToken(userUUID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (this *service) GenerateToken(userUUID string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	// secretKey := []byte("S3C123TKEY")

	tokenHourLifespanString := os.Getenv("TOKEN_HOUR_LIFESPAN")
	// tokenHourLifespanString := "24"
	tokenHourLifespan, err := strconv.Atoi(tokenHourLifespanString)
	if err != nil {
		return "", err
	}

	claim := jwt.MapClaims{}
	claim["user_uuid"] = userUUID
	claim["exp"] = time.Now().Add(time.Hour * time.Duration(tokenHourLifespan)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenGenerated, err := token.SignedString(secretKey)
	if err != nil {
		return tokenGenerated, err
	}

	return tokenGenerated, nil
}

func (this *service) ValidateToken(encodedToken string) (*jwt.Token, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	// secretKey := []byte("S3C123TKEY")

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return secretKey, nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Max-Age", "86400")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(200)
        } else {
            c.Next()
        }
    }
}