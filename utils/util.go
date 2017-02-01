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
func ReturnHTTPErrorResponse(w http.ResponseWriter, httpStatusCode int, errorCode int, message string) {
	responseBody, err := json.Marshal(&ErrorResponse{ErrorCode: errorCode, ErrorMessage: message})
	if err != nil {
		log.Println("Failed to write HTTP Error Response. Error:", err)
	}
	//http.Error(w, string(responseBody[:bytes.IndexByte(responseBody, 0)]), httpStatusCode)
	http.Error(w, string(responseBody), httpStatusCode)
}

var (
	// DBCon is the database connection pool to be used for all data requests
	DBCon *sql.DB
)
