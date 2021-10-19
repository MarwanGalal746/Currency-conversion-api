package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"log"
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

func isAuthorized(APIkey string) bool {
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

func GetAPIkey(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp := make(map[string]string)
	apiKey, _ := GenerateAPIkey()
	resp["API_KEY"]= apiKey
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}



func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/apikey", GetAPIkey).Methods("GET")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}