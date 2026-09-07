package main

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID returns a random 8-character hex string used as a paste's ID.
func GenerateID() (string, error) {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
