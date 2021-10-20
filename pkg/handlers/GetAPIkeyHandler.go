package handlers

import (
	"Currency_conversion_API/pkg/utils"
	"encoding/json"
	"log"
	"net/http"
)

func GetAPIkey(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	resp := make(map[string]string)
	apiKey, _ := utils.GenerateAPIkey()
	resp["API_KEY"]= apiKey
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
