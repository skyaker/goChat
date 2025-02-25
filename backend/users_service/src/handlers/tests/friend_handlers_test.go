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
		AcceptorId: 3,
		Aim:        friend_handlers.SendRequest,
	}
	err := testStruct.SendFriendRequest(db)
	assert.NoError(t, err)
}

func TestDeleteRequest(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   1,
		AcceptorId: 3,
		Aim:        friend_handlers.DeleteRequest,
	}
	err := testStruct.DeleteFriendRequest(db)
	assert.NoError(t, err)
}

func TestAcceptRequest(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   1,
		AcceptorId: 3,
		Aim:        friend_handlers.Accept,
	}
	err := testStruct.AcceptRequest(db)
	assert.NoError(t, err)
}

func TestRejectRequest(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   2,
		AcceptorId: 1,
		Aim:        friend_handlers.Reject,
	}
	err := testStruct.RejectRequest(db)
	assert.NoError(t, err)
}

func TestDeleteFriend(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   3,
		AcceptorId: 1,
		Aim:        friend_handlers.Delete,
	}
	err := testStruct.DeleteFriend(db)
	assert.NoError(t, err)
}

func TestBlockUser(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   6,
		AcceptorId: 5,
		Aim:        friend_handlers.Block,
	}
	err := testStruct.BlockUser(db)
	assert.NoError(t, err)
}

func TestUnblockUesr(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testStruct := friend_handlers.RequestInfo{
		SenderId:   6,
		AcceptorId: 5,
		Aim:        friend_handlers.Unblock,
	}
	err := testStruct.UnblockUser(db)
	assert.NoError(t, err)
}
