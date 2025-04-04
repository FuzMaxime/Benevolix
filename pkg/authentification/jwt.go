package authentification

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(secret string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 999999).Unix(),
	})
	return token.SignedString([]byte(secret))
}

func ParseToken(secret, tokenString string) (uint, error) {
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(claims["userId"].(float64)), nil
	}
	return 0, err
}
