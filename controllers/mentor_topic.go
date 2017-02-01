package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/TheBeege/mentor-me/models"
	"github.com/TheBeege/mentor-me/utils"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// NewMentorTopic swagger:route POST /api/v1/mentor_topic Mentors NewMentorTopic
//
// Constructs and persists a new MentorTopic based on the provided parameters
//
// Responses:
//    default: errorResponse
//        200: models.MentorTopic
func NewMentorTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var mt models.MentorTopic
	err := decoder.Decode(&mt)
	if err != nil {
		log.Println("Error decoding request body for new MentorTopic. Body:", r.Body, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, http.StatusBadRequest, 200, "Request body was not properly formatted")
		return
	}

	tx, err := utils.DBCon.Begin()
	if err != nil {
		log.Println("Error starting new transaction for inserting mentor topic. Error:", err, "-- MentorTopic:", mt)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 210, "Error creating new mentor topic")
		return
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
    INSERT INTO main.mentor_topic
    (
			user_id
			,topic_id
			,level
			,description
		)
    VALUES ($1, $2, $3, $4)
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

// FindMentors swagger:route GET /api/v1/find_mentors/{topic}/{level} Mentors FindMentors
//
// Returns a list of mentors that are listed for the given topic at or above the given level
//
// Responses:
//    default: errorResponse
//        200: []models.User
func FindMentors(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)
	topicName, ok := urlVars["topic"]
	if !ok {
		log.Println("Error retrieving topic name from request. urlVars:", urlVars)
		utils.ReturnHTTPErrorResponse(w, http.StatusBadRequest, 300, "Topic name missing from request")
		return
	}
	level, err := strconv.Atoi(urlVars["level"])
	if err != nil {
		log.Println("Error retrieving level from request. urlVars:", urlVars, "Error:", err)
		utils.ReturnHTTPErrorResponse(w, http.StatusBadRequest, 300, "Level missing from request")
		return
	}

	userSlice := make([]*models.User, 0)
	rows, err := utils.DBCon.Query(`
	SELECT
		u.id
		,u.username
		,u.display_name
		,u.created
		,u.last_activity
		,u.description
		,u.icon_url
	FROM main.mentor_topic mt
	JOIN main.user u
	ON mt.user_id = u.id
	JOIN main.topic t
	ON mt.topic_id = t.id
	WHERE t.name = $1
	AND mt.level >= $2
	`, topicName, level)
	if err != nil {
		log.Println("Error retrieving mentors from database. TopicName:", topicName, "Level:", level, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 320, "Error retrieving mentors")
		return
	}
	for rows.Next() {
		user := new(models.User)
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.DisplayName,
			&user.Created,
			&user.LastActivity,
			&user.Description,
			&user.IconURL,
		); err != nil {
			log.Println("Error scanning mentor user from database. TopicName:", topicName, "Level:", level, "-- Error:", err)
			utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 310, "Error retrieving mentors")
			continue
		}
		userSlice = append(userSlice, user)
	}

	if err := json.NewEncoder(w).Encode(userSlice); err != nil {
		log.Println("Error encoding response for requested topics. UsersSlice:", userSlice, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, http.StatusInternalServerError, 340, "Error retrieving similar topic")
		return
	}
}
