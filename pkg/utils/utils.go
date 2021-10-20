package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
)

var key = []byte("currencyTask")


//GenerateAPIkey to generate API key
func GenerateAPIkey() (string, error){
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	min := 0
	max := 987654321
	claims["rand"] = rand.Intn(max - min + 1) + min
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenString, nil
}

func IsAuthorized(APIkey string) bool {
	token, err := jwt.Parse(APIkey, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return key, nil
	})
	if err != nil {
		return false
	}
	if token.Valid{
		return true
	} else {
		return false
	}
}

