package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"postgressql/models"
)

func List(w http.ResponseWriter, r *http.Request){
	todo, err := models.GetAll()
	if err != nil {
		log.Printf("Erro : %v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}