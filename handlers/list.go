package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgressql/models"
)

func List(w http.ResponseWriter, r *http.Request){
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro : %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(w).Encode(todos)
}