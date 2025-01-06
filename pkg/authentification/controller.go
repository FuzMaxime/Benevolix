package authentification

import (
	"benevolix/config"
	"encoding/json"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"

	_ "benevolix/docs" // Importez les fichiers de documentation générés
)

type LoginConfig struct {
	*config.Config
}

func New(configuration *config.Config) *LoginConfig {
	return &LoginConfig{configuration}
}

// LoginPayload représente le payload pour la connexion
type LoginPayload struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// @Summary Connexion de l'utilisateur
// @Description Permet à un utilisateur de se connecter avec ses identifiants.
// @Tags Authentification
// @Accept json
// @Produce json
// @Param payload body LoginPayload true "Login payload"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]
// {
// 	"email": example@example.com"
// 	"password": P@ssw0rd"
// }
func (config *LoginConfig) Login(w http.ResponseWriter, r *http.Request) {
	var payload LoginPayload
	
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
