package handlers

import (
	"database/sql"
	"fmt"
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

func (m *MessageCreate) CheckSender(db *sql.DB) error {
	var userId1, userId2 uint

	query := `SELECT user_1_id, user_2_id
						FROM dialogs
						WHERE id = $1`

	rows, err := db.Query(query, m.DialogId)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&userId1, &userId2); err != nil {
			return fmt.Errorf("failed to scan users from dialog with id: %v", m.DialogId)
		}
	}

	if m.SenderId != userId1 && m.SenderId != userId2 {
		return fmt.Errorf("user was not found in dialog")
	} else {
		return nil
	}
}

func (m *MessageCreate) SendMessage(db *sql.DB) error {
	if err := m.CheckSender(db); err != nil {
		return err
	}
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
