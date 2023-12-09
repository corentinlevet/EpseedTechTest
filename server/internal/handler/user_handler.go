package handler

import (
	"encoding/json"
	"epseed/internal/db"
	"net/http"
)

// GET /users/get
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
		return
	}

	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, "Erreur de requête des utilisateurs", http.StatusInternalServerError)
		return
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Erreur d'encodage JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(usersJSON)
}
