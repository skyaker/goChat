package handlers

import (
	"database/sql"
	"time"
)

// TODO: last message update

type MessageCreate struct {
	DialogId  uint      `json:"dialog_id"`
	SenderId  uint      `json:"sender_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type MessageDelete struct {
	MessageId uint `json:"message_id"`
}

type MessageEdit struct {
	MessageId  uint   `json:"message_id"`
	NewContent string `json:"content"`
}

func (m *MessageCreate) SendMessage(db *sql.DB) error {
	query := `INSERT INTO messages (dialog_id, sender_id, content, created_at) 
		  			VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, m.DialogId, m.SenderId, m.Content, m.CreatedAt)
	return err
}

func (m *MessageDelete) DeleteMessage(db *sql.DB) error {
	query := `DELETE FROM messages 
						WHERE message_id = $1;`
	_, err := db.Exec(query, m.MessageId)
	return err
}

func (m *MessageEdit) EditMessage(db *sql.DB) error {
	query := `UPDATE messages
						SET content = $1
						WHERE message_id = $2`
	_, err := db.Exec(query, m.NewContent, m.MessageId)
	return err
}
