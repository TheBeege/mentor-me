package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/TheBeege/mentor-me/models"
	"github.com/TheBeege/mentor-me/utils"
	"github.com/lib/pq"
)

// NewMentorRequest swagger:route POST /api/v1/mentor_request Mentors NewMentorRequest
//
// Constructs and persists a new MentorRequest based on the provided parameters
//
// Responses:
//    default: errorResponse
//        200: models.MentorRequest
func NewMentorRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var mt models.MentorTopic
	err := decoder.Decode(&mt)
	if err != nil {
		log.Println("Error decoding request body for new MentorRequest. Body:", r.Body, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, http.StatusBadRequest, 200, "Request body was not properly formatted")
		return
	}

	tx, err := utils.DBCon.Begin()
	if err != nil {
		log.Println("Error starting new transaction for inserting mentor request. Error:", err, "-- MentorRequest:", mt)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 210, "Error creating new mentor request")
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
    INSERT INTO main.mentor_request
    (
			mentor_id
			,mentee_id
			,requested
		)
    VALUES ($1, $2, $3)
  `, mt.UserID, mt.TopicID, mt.Level, mt.Description)
	if err != nil {
		if err.(*pq.Error).Code.Name() == "unique_violation" {
			constraintPattern, _ := regexp.Compile("violates unique constraint \"(.+?)\"")
			matchSlice := constraintPattern.FindSubmatch([]byte(err.(*pq.Error).Message))
			if len(matchSlice) < 2 {
				log.Println("Received request to create mentor topic but encountered unknown unique constraint violation. Error:", err)
				utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 230, "Unknown error attempting to create mentor topic")
				return
			}
			constraint := string(matchSlice[1])
			if constraint == "mentor_topic_id" {
				log.Println("Received request to create mentor topic with duplicate user and topic IDs. UserID:", mt.UserID, "TopicID:", mt.TopicID)
				utils.ReturnHTTPErrorResponse(w, http.StatusBadRequest, 240, "Mentor topic combination already in use")
				return
			}
			// TODO: Is there a better way to organize this?
		}
		log.Println("Error inserting new mentor topic. Error:", err, "-- MentorTopic:", mt)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 260, "Error creating new mentor topic")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction for new mentor topic. Error:", err, "-- MentorTopic:", mt)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 280, "Error creating new mentor topic")
		return
	}
	log.Println("Successfully created new mentor topic")
}

// AcceptMentorRequest swagger:route POST /api/v1/mentor_request/{id:[0-9]}/accept Mentors AcceptMentorRequest
//
// Marks a MentorRequest as accepted
//
// Responses:
//    default: errorResponse
//        200: models.MentorRequest
func AcceptMentorRequest(w http.ResponseWriter, r *http.Request) {

}

// RejectMentorRequest swagger:route POST /api/v1/mentor_request/{id:[0-9]}/reject Mentors RejectMentorRequest
//
// Constructs and persists a new MentorRequest based on the provided parameters
//
// Responses:
//    default: errorResponse
//        200: models.MentorRequest
func RejectMentorRequest(w http.ResponseWriter, r *http.Request) {

}

// GetMentorRequestList swagger:route GET /api/v1/mentor_request/mentor/{mentor_id:[0-9]} Mentors GetMentorRequests
//
// Retrieves all MentorRequests for a given mentor ID
//
// Responses:
//    default: errorResponse
//        200: models.MentorRequest
func GetMentorRequestList(w http.ResponseWriter, r *http.Request) {

}

// GetMentorRequest swagger:route GET /api/v1/mentor_request/{id:[0-9]} Mentors GetMentorRequests
//
// Retrieves information about a MentorRequest for a given ID
//
// Responses:
//    default: errorResponse
//        200: models.MentorRequest
func GetMentorRequest(w http.ResponseWriter, r *http.Request) {

}
