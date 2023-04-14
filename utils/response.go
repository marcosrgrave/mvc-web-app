package utils

import (
	"encoding/json"
	"net/http"
)

func JSONResponseMessage(w http.ResponseWriter, statusCode int, msg string) {
	msgResponse := map[string]string{"msg": msg}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msgResponse)
}

func JSONResponseError(w http.ResponseWriter, statusCode int, msg string, err error) {
	msgResponse := map[string]string{"msg": msg + err.Error()}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(msgResponse)
}

func JSONResponseData(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
