package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"postgressql/models"
)

func Create(w http.ResponseWriter, r *http.Request){
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer Decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)
	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Erro na query"),
		}
	}else{
		resp = map[string]any{
			"Error": false,
			"Message": fmt.Sprintf("Inserido com sucesso ID: %d", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}