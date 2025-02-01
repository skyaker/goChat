package handlers_test

import (
	"database/sql"
	"fmt"
	dbconn "messages_service/db_connection"
	"messages_service/handlers"
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

	testDialog := handlers.NewDialog{
		UserId_1:        1,
		UserId_2:        2,
		LastMessage:     "Hello David",
		LastMessageTime: time.Now(),
	}

	err := testDialog.CreateDialog(db)
	assert.NoError(t, err)
}
