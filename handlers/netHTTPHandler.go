package handlers

import (
	"encoding/json"
	"net/http"
)

// HTTPLoginHandler sets up the login handler using the standard net/http library
func HTTPLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	resp := ValidateLoginPost(r.Body)

	json.NewEncoder(w).Encode(resp)
}
