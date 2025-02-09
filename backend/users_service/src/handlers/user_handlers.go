package user_handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

// UserCreateInfo info structure for sign up
type UserCreateInfo struct {
	UserId      uint      `json:"user_id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"created_at"`
}

// UserTokenBody required info for getting token
type UserTokenBody struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
}

// UserLoginInfo required info for logging in
type UserLoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserDeleteInfo info structure for removing
type UserDeleteInfo struct {
	Id uint `json:"id"`
}

type NewUsername struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type NewPassword struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}

type NewEmail struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

func checkUserExistanceToInsert(db *sql.DB, username *string, email *string) error {
	query := `SELECT id 
						FROM users 
						WHERE username = $1 OR email = $2`
	row := db.QueryRow(query, username, email)
	err := row.Scan(new(int))

	return err
}

func checkUserLoginData(db *sql.DB, username *string, password *string) error {
	var realPassword string
	query := `SELECT password
						FROM users
						WHERE username = $1`
	err := db.QueryRow(query, username).Scan(&realPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		} else {
			return err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(*password))
	if err != nil {
		return fmt.Errorf("invalid username or password")
	}

	return nil
}

func checkUserExistanceToDelete(db *sql.DB, id *uint) error {
	query := `SELECT * 
						FROM users
						WHERE id = $1`
	row := db.QueryRow(query, id)
	err := row.Scan(new(int))

	return err
}

func getUserToken(userId *uint, username *string) (string, error) {
	var url string = "http://localhost:8081/auth/token"
	var info = UserTokenBody{
		UserId:   *userId,
		Username: *username,
	}

	jsonData, err := json.Marshal(info)
	if err != nil {
		return "", fmt.Errorf("error encoding JSON: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("auth service returned status: %d", err)
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	token, status := result["token"]
	if !status {
		return "", fmt.Errorf("token was not found in response")
	}

	return token, nil
}

// AddUser New user creation
// @Description Inserts user in db
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserCreateInfo true "User data"
// @Success 201 {object} map[string]string "User registered successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 409 {object} map[string]string "User already exists"
// @Failure 500 {object} map[string]string "Database error"
// @Router /register [post]
func AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData UserCreateInfo
		var exists bool = true

		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err := checkUserExistanceToInsert(db, &userData.Username, &userData.Email)

		if err != nil {
			if err == sql.ErrNoRows {
				exists = false
			} else {
				http.Error(w, "Database error", http.StatusConflict)
				return
			}
		}

		if exists {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		query := `INSERT INTO users (username, password, email, created_at)
							VALUES ($1, $2, $3, $4) RETURNING id`
		err = db.QueryRow(query, userData.Username, string(hashedPassword), userData.Email, userData.DateCreated).Scan(&userData.UserId)

		if err != nil {
			http.Error(w, "Database insert error", http.StatusInternalServerError)
			return
		}

		userToken, err := getUserToken(&userData.UserId, &userData.Username)
		if err != nil {
			http.Error(w, "Error getting token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User registered successfully",
			"token":   userToken,
		})
	}
}

// Login User logging in
// @Description Returns token to user after login, password confirmation
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserLoginInfo true "User data"
// @Success 201 {object} map[string]string "User registered successfully"
// @Failure 400 {Object} map[string]string "Invalid request"
// @Failure 401 {Object} map[string]string "Invalid username or password"
// @Failure 500 {object} map[string]string "Database error"
// @Router /login [post]
func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userData UserLoginInfo
		var userId uint

		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		err := checkUserLoginData(db, &userData.Username, &userData.Password)

		if err != nil {
			if err.Error() == "invalid username or password" {
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			} else {
				http.Error(w, "Database error", http.StatusInternalServerError)
			}
			return
		}

		query := `SELECT id
							FROM users
							WHERE username = $1`
		err = db.QueryRow(query, userData.Username).Scan(&userId)

		if err != nil {
			http.Error(w, "Error getting user id", http.StatusInternalServerError)
			return
		}

		userToken, err := getUserToken(&userId, &userData.Username)
		if err != nil {
			http.Error(w, "Error getting token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "User registered successfully",
			"token":   userToken,
		})
	}
}

// DeleteUser User removal
// @Description Deletes user from db
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserDeleteInfo true "Delete data"
// @Success 204 {object} map[string]string "User was deleted successfully"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 404 {object} map[string]string "User doesn't exist"
// @Failure 500 {object} map[string]string "Database error"
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var delData UserDeleteInfo
		idParam := chi.URLParam(r, "id")

		userID64, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		userID := uint(userID64)

		err = checkUserExistanceToDelete(db, &userID)

		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "Cannot delete: user does not exist", http.StatusNotFound)
				return
			}
			http.Error(w, "Database error", http.StatusConflict)
			return
		}

		query := `DELETE FROM users
							WHERE id = $1`
		_, err = db.Exec(query, delData.Id)

		if err != nil {
			http.Error(w, "Database delete error", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (newUsernameData *NewUsername) ChangeUsername(db *sql.DB) error {
	query := `UPDATE users
						SET username = $2
						WHERE id = $1`

	_, err := db.Exec(query, newUsernameData.Id, newUsernameData.Username)
	return err
}

func (newPasswordData *NewPassword) ChangePassword(db *sql.DB) error {
	query := `UPDATE users
						SET password = $2
						WHERE id = $1`

	_, err := db.Exec(query, newPasswordData.Id, newPasswordData.Password)
	return err
}

func (newEmailData *NewEmail) ChangeEmail(db *sql.DB) error {
	query := `UPDATE users
						SET email = $2
						WHERE id = $1`

	_, err := db.Exec(query, newEmailData.Id, newEmailData.Email)
	return err
}
