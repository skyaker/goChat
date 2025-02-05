package user_handlers_test

import (
	"database/sql"
	"fmt"
	"testing"
	dbconn "users_service/database/db_connection"
	friend_handlers "users_service/src/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func TestFriendRequest(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   1,
		AcceptorId: 5,
		Aim:        friend_handlers.SendRequest,
	}
	err := testStruct.FriendRequest(db)
	assert.NoError(t, err)
}
