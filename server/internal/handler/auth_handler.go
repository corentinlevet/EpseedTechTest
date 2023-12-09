package handler

import (
	"encoding/json"
	"epseed/internal/db"
	"net/http"
)

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	UserId  int    `json:"user_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func writeErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	response := AuthResponse{
		Message: message,
		Token:   "",
	}

	returnJson, _ := json.Marshal(response)
	w.Write(returnJson)
}

// POST /auth/login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		writeErrorResponse(w, "Erreur de lecture du formulaire")
		return
	}

	user, err := db.GetUserByUsernameAndPassword(loginRequest.Username, loginRequest.Password)
	if err != nil {
		writeErrorResponse(w, "Erreur de requête des utilisateurs")
		return
	}

	if user == nil {
		writeErrorResponse(w, "Utilisateur non trouvé")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(AuthResponse{
		Message: "Utilisateur connecté avec succès",
		Token:   "token",
		UserId:  int(user.ID),
	})

	w.Write(returnJson)
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /auth/signup
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Méthode HTTP non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var signupRequest SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signupRequest)
	if err != nil {
		writeErrorResponse(w, "Erreur de lecture du formulaire")
		return
	}

	user, err := db.GetUserByEmail(signupRequest.Email)
	if err != nil {
		writeErrorResponse(w, "Erreur de requête des utilisateurs")
		return
	}

	if user != nil {
		writeErrorResponse(w, "Un utilisateur avec cet email existe déjà")
		return
	}

	err = db.CreateUser(signupRequest.Username, signupRequest.Email, signupRequest.Password)
	if err != nil {
		writeErrorResponse(w, "Erreur de création de l'utilisateur")
		return
	}

	newlyCreatedUser, err := db.GetUserByEmail(signupRequest.Email)
	if err != nil {
		writeErrorResponse(w, "Erreur de requête des utilisateurs")
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Add("Content-Type", "application/json")

	var returnJson, _ = json.Marshal(AuthResponse{
		Message: "Utilisateur créé avec succès",
		Token:   "token",
		UserId:  int(newlyCreatedUser.ID),
	})

	w.Write(returnJson)
}
