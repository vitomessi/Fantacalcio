package utils

import (
	"encoding/json"
	"net/http"
)
//Utile per visualizzare i dati in formato JSON
func ResponseJSON(w http.ResponseWriter, p interface{}, status int){
	encodeByte, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		http.Error(w, "error om", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(encodeByte))
}
