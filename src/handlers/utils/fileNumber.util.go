package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// GenerateUniqueFileNumber generates a unique and powerful number for files in the system.
func GenerateUniqueFileNumber() (string, error) {
	// Generate a random 4-byte slice
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Convert the random bytes to hexadecimal string
	randomHex := hex.EncodeToString(randomBytes)

	// Get the current timestamp in milliseconds
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	// Concatenate the timestamp and random hex to create the unique file number
	uniqueNumber := fmt.Sprintf("%d%s", timestamp, randomHex)

	return uniqueNumber, nil
}
