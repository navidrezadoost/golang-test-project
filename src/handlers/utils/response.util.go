package utils

import (
	"time"
)

// ResponseCorrectInterface represents the structure of a correct HTTP response.
type ResponseCorrectInterface struct {
	Status   string                 `json:"status"`
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Error    map[string]interface{} `json:"error,omitempty"`
	Data     interface{}            `json:"data,omitempty"`
	Metadata Metadata               `json:"metadata"`
}

// Metadata represents the metadata associated with an HTTP response.
type Metadata struct {
	Version   string    `json:"version"`
	LogID     string    `json:"log_id"`
	Timestamp time.Time `json:"timestamp"`
}

// ResponseCorrect generates a correct HTTP response object with the provided data and status code.
func ResponseCorrect(data interface{}, code int, loggerID string, status string, message string) ResponseCorrectInterface {
	response := ResponseCorrectInterface{
		Status:  status,
		Code:    code,
		Message: message,
		Metadata: Metadata{
			Version:   "1.0.0", // Hardcoded version since there's no package.json in Go
			LogID:     loggerID,
			Timestamp: time.Now(),
		},
	}

	// Modify the response object based on the status
	if status == "error" {
		response.Error = data.(map[string]interface{})
	} else {
		response.Data = data
	}

	return response
}
