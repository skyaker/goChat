package user_handlers

import (
	"database/sql"
	"fmt"
)

type FriendRequestInfo struct {
	SenderId   uint `json:"user_id"`
	AcceptorId uint `json:"friend_id"`
}

func (reqInfo *FriendRequestInfo) checkFriendStatus(db *sql.DB) (string, error) {
	var status string
	query := `SELECT status 
						FROM friends
						WHERE user_id = $1 AND friend_id = $2`
	err := db.QueryRow(query, reqInfo.SenderId, reqInfo.AcceptorId).Scan(&status)

	if err == sql.ErrNoRows {
		return "free", nil
	}

	if err != nil {
		return "", err
	} else {
		return status, nil
	}
}

func (reqInfo *FriendRequestInfo) FriendRequest(db *sql.DB) error {
	status, err := reqInfo.checkFriendStatus(db)

	switch status {
	case "free":
		query := `INSERT INTO friends (user_id, friend_id, status)
		VALUES ($1, $2, $3)`

		_, err = db.Exec(query, reqInfo.SenderId, reqInfo.AcceptorId, "pending")
		return err
	case "accepted":
		return fmt.Errorf("users are already friends")
	case "pending":
		return fmt.Errorf("request was already sent")
	case "blocked":
		return fmt.Errorf("trying to add blocked user")
	default:
		return err
	}
}

// func AcceptFriendRequest

// func DeleteFriend

// func BlockUser

// func Unblock (to do free)
