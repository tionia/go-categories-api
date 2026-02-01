package handlers

import (
	"encoding/json"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Server is running",
	})
}
