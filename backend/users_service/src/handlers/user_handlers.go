package user_handlers

import (
	"database/sql"
	"time"
)

// добавление пользователя
// удаление пользователя
// смена username
// смена пароля
// смена почты
// смена аватарки (?)

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

// func checkUserExistance(db *sql.DB, username *string, email *string) error {
// 	query := `SELECT * FROM users
// 						WHERE username = $1 and email = $2`
// 	rows, err := db.Query(query, username, email)

// 	if err != nil {
// 		return err
// 	}

// 	if rows.Next() {
// 		if err = rows.Scan(&username); err == nil {
// 			return fmt.Errorf("username %v alredy exists", username)
// 		} else if err = rows.Scan(&email); err == nil {
// 			return fmt.Errorf("account with email \"%v\" alredy exists", email)
// 		}
// 	}

// 	return nil
// }

func (userData *UserCreateInfo) AddUser(db *sql.DB) error {
	query := `INSERT INTO users (username, password, email, created_at)
						VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, userData.Username, userData.Password, userData.Email, userData.DateCreated)
	return err
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
