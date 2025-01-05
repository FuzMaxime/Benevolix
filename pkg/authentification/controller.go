package authentification

import (
	"benevolix/config"
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type LoginConfig struct {
	*config.Config
}

func New(configuration *config.Config) *LoginConfig {
	return &LoginConfig{configuration}
}

func (config *LoginConfig) Login(w http.ResponseWriter, r *http.Request) {
	print("test")
	var payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&payload)
	user, exist := config.UserRepository.GetUserByEmail(payload.Email)

	if exist != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)) != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateToken(os.Getenv("API_Key"), payload.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
