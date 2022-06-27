package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"regexp"
	"rehber/models"
	"strconv"
)

//IsPhoneValid is a function which checks if the phone number is valid.
//If phone number is not valid, it will return false and the index of the invalid phone number.
//If phone number is valid, it will return true and -1.
func IsPhoneValid(phones []models.PhoneNumber) (bool, int) {
	r, _ := regexp.Compile(`^(0|\+62|062|62)[0-9]+$`)

	for i, v := range phones {
		if len(v.Number) < 11 {
			return false, i
		}
		if !r.MatchString(v.Number) {
			return false, i
		}
	}
	return true, -1
}

//IsTokenValid is a function which checks if the JWT is valid.
func IsTokenValid(id string, t *jwt.Token) bool {
	i, _ := strconv.Atoi(id)
	tid := int(t.Claims.(jwt.MapClaims)["user_id"].(float64))

	if i != tid {
		return false
	}
	return true
}
