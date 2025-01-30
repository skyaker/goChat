package handlers

import (
	"database/sql"
	"time"
)

// TODO: last message update

type NewMessage struct {
	DialogId  uint      `json:"dialog_id"`
	SenderId  uint      `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type DeletedMessage struct {
	MessageId uint `json:"message_id"`
}

type EditedMessage struct {
	MessageId  uint   `json:"message_id"`
	NewContent string `json:"content"`
}

func (m *NewMessage) SendMessage(db *sql.DB) {
	query := `INSERT INTO messages (dialog_id, sender_id, content, created_at) 
		  			VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, m.DialogId, m.SenderId, m.Content, m.CreatedAt)
	if err != nil {
		panic(err)
	}
}

func (m *DeletedMessage) DeleteMessage(db *sql.DB) {
	query := `DELETE FROM messages 
						WHERE message_id = $1;`
	_, err := db.Exec(query, m.MessageId)
	if err != nil {
		panic(err)
	}
}

func (m *EditedMessage) EditMessage(db *sql.DB) {
	query := `UPDATE messages
						SET content = $1
						WHERE message_id = $2`
	_, err := db.Exec(query, m.NewContent, m.MessageId)
	if err != nil {
		panic(err)
	}
}
