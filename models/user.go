package models

import (
	"simple-login/db"

	"github.com/google/uuid"
)

func ValidateUser(username, password string) bool {
	var dbPass string
	err := db.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&dbPass)
	if err != nil || dbPass != password {
		return false
	}
	return true
}

func CreateSession(username string) string {
	sessionID := uuid.New().String()
	db.DB.Exec("INSERT INTO sessions(id, username) VALUES (?, ?)", sessionID, username)
	return sessionID
}

func GetUsernameFromSession(sessionID string) string {
	var username string
	err := db.DB.QueryRow("SELECT username FROM sessions WHERE id = ?", sessionID).Scan(&username)
	if err != nil {
		return ""
	}
	return username
}

func DeleteSession(sessionID string) {
	db.DB.Exec("DELETE FROM sessions WHERE id = ?", sessionID)
}
