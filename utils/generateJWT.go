package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"rehber/config"
	"time"
)

//GenerateToken is a function which create a JWT.
func GenerateToken(id uint) string {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(time.Minute * 30).Unix(),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := t.SignedString([]byte(config.Config("JWT_SECRET")))
	return token
}
