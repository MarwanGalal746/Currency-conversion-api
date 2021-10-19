package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"net/http"
	"time"
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



func main() {
	rand.Seed(time.Now().UnixNano())
	str, _ := GenerateAPIkey()
	fmt.Println(str)
}