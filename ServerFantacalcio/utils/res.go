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
	//setta come sarà il formato
	w.Header().Set("Content-Type", "application/json")
	//setta lo status, ovvero il codice
	w.WriteHeader(status)
	//aggiunge i dati, in questo caso quelli codificati, ad una response
	w.Write([]byte(encodeByte))
}
