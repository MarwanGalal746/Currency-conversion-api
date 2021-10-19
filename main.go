package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

type currency struct {
	Success    bool   `json:"success"`
	Timestamp  int    `json:"timestamp"`
	Historical bool   `json:"historical"`
	Base       string `json:"base"`
	Date       string `json:"date"`
	Rates      struct {
	USD float64 `json:"USD"`
	} `json:"rates"`
}

type Response struct {
	Message string `json:"message"`
	Status int `json:"status"`
}
func NewResponse(message string, status int) *Response {
	return &Response{Message: message, Status: status}
}

func GetCurrency(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	api_key, _ := vars["API_key"]
	if !isAuthorized(api_key) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(NewResponse("Invalid API key", http.StatusUnauthorized))
		return
	}
	resp, err := http.Get("http://data.fixer.io/api/2020-10-19?access_key=d2f10c93785cdb3dfd738ad60dca039d&symbols=USD&format=1")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var curr currency
	err = json.Unmarshal(body, &curr)
	response := make(map[string]float64)
	response["currency rate from EUR to USD"]= curr.Rates.USD
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func ConvertUSDtoEUR(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	api_key, _ := vars["API_key"]
	usd, _ := vars["usd"]
	usdFloat, _ := strconv.ParseFloat(usd,64)
	fmt.Println(usdFloat)
	if !isAuthorized(api_key) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(NewResponse("Invalid API key", http.StatusUnauthorized))
		return
	}
	resp, err := http.Get("http://data.fixer.io/api/2020-10-19?access_key=d2f10c93785cdb3dfd738ad60dca039d&symbols=USD&format=1")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var curr currency
	err = json.Unmarshal(body, &curr)
	response := make(map[string]float64)
	response["Amount in EUR equals"]= usdFloat / curr.Rates.USD
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

func ConvertEURtoUSD(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	api_key, _ := vars["API_key"]
	eur, _ := vars["eur"]
	eurFloat, _ := strconv.ParseFloat(eur,64)
	fmt.Println(eurFloat)
	if !isAuthorized(api_key) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(NewResponse("Invalid API key", http.StatusUnauthorized))
		return
	}
	resp, err := http.Get("http://data.fixer.io/api/2020-10-19?access_key=d2f10c93785cdb3dfd738ad60dca039d&symbols=USD&format=1")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var curr currency
	err = json.Unmarshal(body, &curr)
	response := make(map[string]float64)
	response["Amount in eur equals"]= eurFloat * curr.Rates.USD
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}



func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/apikey", GetAPIkey).Methods("GET")
	r.HandleFunc("/currency_rate", GetCurrency).Methods("GET").Queries("api_key" , "{API_key}")
	r.HandleFunc("/convert_usd", ConvertUSDtoEUR).Methods("GET").Queries("api_key" , "{API_key}" , "usd", "{usd}")
	r.HandleFunc("/convert_eur", ConvertEURtoUSD).Methods("GET").Queries("api_key" , "{API_key}" , "eur", "{eur}")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}