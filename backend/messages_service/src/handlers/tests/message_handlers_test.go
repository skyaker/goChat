package handlers_test

import (
	dbconn "messages_service/database/db_connection"

	"database/sql"
	message_handlers "messages_service/src/handlers"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestMessageDispatch(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testMessage := message_handlers.MessageCreate{
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

	testDelMessage := message_handlers.MessageDelete{
		MessageId: 22,
	}

	err := testDelMessage.DeleteMessage(db)
	assert.NoError(t, err)
}

func TestEditMessage(t *testing.T) {
	var db *sql.DB = dbconn.GetDbConnection()

	testEdMessage := message_handlers.MessageEdit{
		MessageId:  21,
		NewContent: "Hello John",
	}

	err := testEdMessage.EditMessage(db)
	assert.NoError(t, err)
}
