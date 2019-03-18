package utils

import (
	"encoding/json"
	"net/http"
)

// Message package a message.
func Message(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}

// Response create a HTTP response.
func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// NewError create a HTTP error response.
func NewError(err error, w http.ResponseWriter) {
	res := Message("error")
	res["data"] = err
	Response(w, res)
}

// CheckError check if an error occur and if there's an error, return an error HTTP response.
func CheckError(err error, w http.ResponseWriter) bool {
	if err != nil {
		NewError(err, w)
		return true
	}
	return false
}
