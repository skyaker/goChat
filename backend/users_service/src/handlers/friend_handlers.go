package user_handlers

import (
	"database/sql"
	"fmt"
)

type Purpose string

const (
	SendRequest    Purpose = "sendRequest"
	DeclineRequest Purpose = "declineRequest"
	Accept         Purpose = "accept"
	Reject         Purpose = "reject"
	Delete         Purpose = "delete"
	Block          Purpose = "block"
	Unblock        Purpose = "unblock"
)

type RequestInfo struct {
	SenderId   uint    `json:"user_1_id"`
	AcceptorId uint    `json:"user_2_id"`
	Aim        Purpose `json:"sender_purpose"`
}

func (reqInfo *RequestInfo) getAccessStatus(db *sql.DB) ([]Purpose, error) {
	var status string
	var status_creator uint

	query := `SELECT status, status_creator
						FROM relations
						WHERE (user_1_id = $1 AND user_2_id = $2) 
							 OR (user_1_id = $2 AND user_2_id = $1)`

	err := db.QueryRow(query, reqInfo.SenderId, reqInfo.AcceptorId).Scan(&status, &status_creator)

	if err == sql.ErrNoRows {
		return []Purpose{SendRequest, Block}, err
	}

	if reqInfo.AcceptorId == status_creator {
		switch status {
		case "pending":
			return []Purpose{DeclineRequest, Block}, nil
		case "accepted":
			return []Purpose{Delete, Block}, nil
		case "blocked":
			return []Purpose{Unblock}, nil
		}
	} else {
		switch status {
		case "pending":
			return []Purpose{Accept, Reject, Block}, nil
		case "accepted":
			return []Purpose{Delete, Block}, nil
		case "blocked":
			return []Purpose{Block}, nil
		}
	}

	return []Purpose{}, err
}

func (reqInfo *RequestInfo) FriendRequest(db *sql.DB) error {
	var query string
	var availableFunctions []Purpose
	var accessible bool = false

	availableFunctions, err := reqInfo.getAccessStatus(db)

	for _, v := range availableFunctions {
		if v == SendRequest {
			accessible = true
		}
	}

	if !accessible {
		return fmt.Errorf("user doesn't have permission for operation: %v", reqInfo.Aim)
	}

	if err == nil {
		query = `UPDATE relations
						 SET user_1_id = $1, user_2_id = $2, status = $3, status_creator = $4
						 WHERE (user_1_id = $1 AND user_2_id = $2) 
							  OR (user_1_id = $2 AND user_2_id = $1)`
	} else if err == sql.ErrNoRows {
		query = `INSERT INTO relations (user_1_id, user_2_id, status, status_creator)
						 VALUES ($1, $2, $3, $4)`
	} else if err != sql.ErrNoRows {
		return err
	}

	_, err = db.Exec(query, reqInfo.SenderId, reqInfo.AcceptorId, "pending", reqInfo.SenderId)
	return err
}
