package handlers

import (
	"database/sql"
	"time"
)

type NewDialog struct {
	UserId_1        uint      `json:"user_1_id"`
	UserId_2        uint      `json:"user_2_id"`
	LastMessage     string    `json:"last_message"`
	LastMessageTime time.Time `json:"last_message_at"`
}

// TODO: add user_id dialog cleared for
type DialogCleaningData struct {
	DialogId uint `json:"dialog_id"`
}

func checkDialogExistanceByUsers(userId_1 *uint, userId_2 *uint, db *sql.DB) error {
	query := `SELECT id FROM dialogs WHERE (user_1_id = $1 AND user_2_id = $2) or (user_1_id = $2 AND user_2_id = $1)`
	_, err := db.Exec(query, userId_1, userId_2)
	return err
}

func checkDialogById(dialogId *uint, db *sql.DB) error {
	query := `SELECT id FROM dialogs WHERE id = $1`
	_, err := db.Exec(query, dialogId)
	return err
}

//TODO: check user existance in users db

func (dialog *NewDialog) CreateDialog(db *sql.DB) error {
	existanceErr := checkDialogExistanceByUsers(&dialog.UserId_1, &dialog.UserId_2, db)
	if existanceErr != nil {
		return existanceErr
	}

	query := `INSERT INTO dialogs (user_1_id, user_2_id, last_message, last_message_at)
	VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, dialog.UserId_1, dialog.UserId_2, dialog.LastMessage, dialog.LastMessageTime)
	return err
}

func (dialog *DialogCleaningData) ClearDialog(db *sql.DB) error {
	existanceErr := checkDialogById(&dialog.DialogId, db)
	if existanceErr != nil {
		return existanceErr
	}

	query := `DELETE FROM messages
						WHERE dialog_id = $1;`
	_, err := db.Exec(query, dialog.DialogId)
	return err
}

func (dialog *DialogCleaningData) DeleteDialog(db *sql.DB) error {
	existanceErr := checkDialogById(&dialog.DialogId, db)
	if existanceErr != nil {
		return existanceErr
	}

	query := `DELETE FROM dialogs
						WHERE id = $1;`
	_, err := db.Exec(query, dialog.DialogId)
	return err
}
