package auth_handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

// UserTokenBody required info for token
type UserTokenBody struct {
	UserId      uint      `json:"user_id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"created_at"`
}

// CreateToken() Creates token
// @Description Creates token based on user info
// @Tags auth
// @Accept json
// @Produce json
// @Param user body UserTokenBody true "User data"
// @Failure 500 {object} map[string]string "Secret key missing or token signing error"
// @Router /auth/token [post]
func CreateToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData UserTokenBody
		key, status := os.LookupEnv("secret_key")

		if !status {
			http.Error(w, "Secret key was not found", http.StatusInternalServerError)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"token":       userData.UserId,
			"username":    userData.Username,
			"email":       userData.Email,
			"dateCreated": userData.DateCreated.Unix(),
			"exp":         time.Now().AddDate(0, 6, 0).Unix(),
		})

		signed, err := token.SignedString([]byte(key))

		if err != nil {
			http.Error(w, "Error signing token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": signed})
	}
}
