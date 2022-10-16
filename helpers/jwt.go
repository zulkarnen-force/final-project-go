package helpers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("rahasia")

func GenerateToken(id uint, email string) string {

	claims := jwt.MapClaims{
		"id": id,
		"email": email,
	}

	fmt.Println("Claims ", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, _ := token.SignedString(secretKey)

	return ss

}


func VerifyToken(c *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in proceed")
	headerToken := c.Request.Header.Get("Authorization")
	isBearer := strings.HasPrefix(headerToken, "Bearer")

	if !isBearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token,) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return secretKey, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}