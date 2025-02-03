package handlers_test

import (
	dbconn "messages_service/database/db_connection"

	"database/sql"
	"fmt"
	message_handlers "messages_service/src/handlers"
	"testing"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func TestDialogCreation(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testDialog := message_handlers.NewDialog{
		UserId_1:        1,
		UserId_2:        4,
		LastMessage:     "Hello David",
		LastMessageTime: time.Now(),
	}

	err := testDialog.CreateDialog(db)
	assert.NoError(t, err)
}

func TestClearDialog(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testCleaningData := message_handlers.DialogCleaningData{
		DialogId: 1,
	}

	err := testCleaningData.ClearDialog(db)
	assert.NoError(t, err)
}

func TestDeleteDialog(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testDeleteData := message_handlers.DialogCleaningData{
		DialogId: 7,
	}

	err := testDeleteData.DeleteDialog(db)
	assert.NoError(t, err)
}
