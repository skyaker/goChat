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
