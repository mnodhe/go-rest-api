package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretkey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	parseToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return nil, err
	}
	tokenIsValid := parseToken.Valid
	if !tokenIsValid {
		return nil, errors.New("token is invalid")
	}
	tokenClaims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	//emailClaim := tokenClaims["email"].(string)
	//userIdClaim := tokenClaims["userId"].(int64)
	//expClaim := tokenClaims["exp"].(int64)
	return tokenClaims, nil
}
