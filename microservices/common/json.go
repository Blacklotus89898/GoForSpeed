package common

import (
	"encoding/json"
	"net/http"
)


func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	
	} 
	return nil
}

func WriteError(w http.ResponseWriter, status int, message string) {

	errorResponse := map[string]string{"error": message}
	WriteJSON(w, status, errorResponse)
	
}