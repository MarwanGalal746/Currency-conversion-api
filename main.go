package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"time"
)



func main() {
	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()
	r.HandleFunc("/apikey", GetAPIkey).Methods("GET")
	r.HandleFunc("/currency_rate", GetCurrency).Methods("GET").Queries("api_key" , "{api_key}")
	r.HandleFunc("/convert_usd", ConvertUSDtoEUR).Methods("GET").Queries("api_key" , "{api_key}" , "usd", "{usd}")
	r.HandleFunc("/convert_eur", ConvertEURtoUSD).Methods("GET").Queries("api_key" , "{api_key}" , "eur", "{eur}")
	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
