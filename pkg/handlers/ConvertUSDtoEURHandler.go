package handlers

import (
	models2 "Currency_conversion_API/pkg/models"
	"Currency_conversion_API/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ConvertUSDtoEUR(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)
	apiKey, _ := vars["api_key"]
	usd, _ := vars["usd"]
	usdFloat, _ := strconv.ParseFloat(usd,64)
	fmt.Println(usdFloat)
	if !utils.IsAuthorized(apiKey) {
		fmt.Println(apiKey)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models2.NewResponse("Invalid API key", http.StatusUnauthorized))
		return
	}
	resp, err := http.Get("http://data.fixer.io/api/2020-10-19?access_key=d2f10c93785cdb3dfd738ad60dca039d&symbols=USD&format=1")
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var curr models2.Currency
	err = json.Unmarshal(body, &curr)
	response := make(map[string]float64)
	response["Amount in EUR equals"]= usdFloat / curr.Rates.USD
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

