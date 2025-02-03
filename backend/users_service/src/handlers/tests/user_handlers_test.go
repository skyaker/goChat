package user_handlers_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"
	dbconn "users_service/database/db_connection"
	user_handlers "users_service/src/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func TestUserInsert(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	userDataTest := user_handlers.UserCreateInfo{
		Username:    "pulse",
		Password:    "1234",
		Email:       "gig@mail.ru",
		DateCreated: time.Now(),
	}

	err := userDataTest.AddUser(db)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	userDataTest := user_handlers.UserDeleteInfo{
		Id: 10,
	}

	err := userDataTest.DeleteUser(db)
	assert.NoError(t, err)
}

func TestUpdateUsername(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	userDataTest := user_handlers.NewUsername{
		Id:       1,
		Username: "UpdatedUser1",
	}

	err := userDataTest.ChangeUsername(db)
	assert.NoError(t, err)
}

func TestUpdatePassword(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	userDataTest := user_handlers.NewPassword{
		Id:       1,
		Password: "UpdatedPassword1",
	}

	err := userDataTest.ChangePassword(db)
	assert.NoError(t, err)
}

func TestUpdateEmail(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	userDataTest := user_handlers.NewEmail{
		Id:    1,
		Email: "UpdatedUser1@example.com",
	}

	err := userDataTest.ChangeEmail(db)
	assert.NoError(t, err)
}
