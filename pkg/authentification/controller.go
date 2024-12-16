package authentification

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// var payload struct {
	// 	Email    string `json:"email"`
	// 	Password string `json:"password"`
	// }
	// json.NewDecoder(r.Body).Decode(&payload)

	// // TODO : check if user and password exist in database

	// hashedPassword, exists := users[payload.Email]
	// if !exists || bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password)) != nil {
	// 	http.Error(w, "Invalid email or password", http.StatusUnauthorized)
	// 	return
	// }

	// token, err := GenerateToken("your_secret_key", payload.Email)
	// if err != nil {
	// 	http.Error(w, "Failed to generate token", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(map[string]string{"token": token})
}
