package routes

import (
	"encoding/json"
	"net/http"

	"github.com/aanchalverma/Machine-Coding/blogging/models"
	"github.com/aanchalverma/Machine-Coding/blogging/storage"
	"github.com/aanchalverma/Machine-Coding/blogging/utils"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	// Retract user payload
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest) // 400
		return
	}

	// Save user details under user/ storage
	if err := storage.SaveUser(user); err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError) // 500
		return
	}
	// If no errors, update the status
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	// Retract username and password for Login
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate user credentials
	user, err := storage.GetUser(credentials.Username)
	if err != nil || user.Password != credentials.Password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate token for the user when they log in
	token, err := utils.GenerateToken(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
