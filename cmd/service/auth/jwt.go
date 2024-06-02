package auth

import (
	"strconv"
	"time"

	"github.com/HarshitPG/ecommerce_api_go/cmd/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second*time.Duration(config.Envs.JWTExpInSec)

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userID": strconv.Itoa(userID),
		"expiredAt":time.Now().Add(expiration).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err!=nil{
		return "",err
	}

	return tokenString,nil
}