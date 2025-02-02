package handlers_test

import (
	"database/sql"
	dbconn "messages_service/db_connection"
	"messages_service/handlers"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMessageDispatch(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testMessage := handlers.MessageCreate{
		DialogId:  1,
		SenderId:  2,
		Content:   "Hello David",
		CreatedAt: time.Now(),
	}

	err := testMessage.SendMessage(db)
	assert.NoError(t, err)
}

func TestDeleteMessage(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testDelMessage := handlers.MessageDelete{
		MessageId: 22,
	}

	err := testDelMessage.DeleteMessage(db)
	assert.NoError(t, err)
}

func TestEditMessage(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testEdMessage := handlers.MessageEdit{
		MessageId:  21,
		NewContent: "Hello John",
	}

	err := testEdMessage.EditMessage(db)
	assert.NoError(t, err)
}
