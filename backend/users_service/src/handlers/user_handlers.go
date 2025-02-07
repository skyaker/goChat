package user_handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserCreateInfo info structure for sign up
type UserCreateInfo struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	DateCreated time.Time `json:"created_at"`
}

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

func checkUserExistance(db *sql.DB, username *string, email *string) (bool, error) {
	query := `SELECT id FROM users 
						WHERE username = $1 OR email = $2`
	row := db.QueryRow(query, username, email)
	err := row.Scan(new(int))

	if err == sql.ErrNoRows {
		return false, nil
	} else if err != sql.ErrNoRows && err != nil {
		return false, err
	} else {
		return true, nil
	}
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
		if err := json.NewDecoder(r.Body).Decode(&userData); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		exists, err := checkUserExistance(db, &userData.Username, &userData.Email)

		if err != nil && err != sql.ErrNoRows {
			http.Error(w, "Database error", http.StatusConflict)
			return
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
							VALUES ($1, $2, $3, $4)`
		_, err = db.Exec(query, userData.Username, string(hashedPassword), userData.Email, userData.DateCreated)

		if err != nil {
			http.Error(w, "Database insert error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
	}
}

func (delData *UserDeleteInfo) DeleteUser(db *sql.DB) error {
	query := `DELETE FROM users
						WHERE id = $1`
	_, err := db.Exec(query, delData.Id)
	return err
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
