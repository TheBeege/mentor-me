package utils

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse is a JSON-formatted reponse for a REST call ending in error
type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// ReturnHTTPErrorResponse is a helper function to construct and return proper
// error responses for REST endpoints
func ReturnHTTPErrorResponse(w http.ResponseWriter, code int, message string) {
	r := &ErrorResponse{ErrorCode: code, ErrorMessage: message}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("Failed to write HTTP Error Response. Error:", err)
	}
}

var (
	// DBCon is the database connection pool to be used for all data requests
	DBCon *sql.DB
)
