package controllers

import (
	"database/sql"
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

// GetTopic swagger:route GET /api/v1/topic/{id} Topics GetTopic
//
// Gets info for a topic.
//
// Responses:
//    default: errorResponse
//        200: models.Topic
func GetTopic(w http.ResponseWriter, r *http.Request) {
	var topicID models.TopicIDParam
	urlVars := mux.Vars(r)
	idString, ok := urlVars["id"]
	if !ok {
		log.Println("Error retrieving topic ID from request. urlVars:", urlVars)
		// TODO: Should we be using similar error numbers or new ones?
		utils.ReturnHTTPErrorResponse(w, 300, "Error fetching topic")
		return
	}

	var err error
	topicID.ID, err = strconv.Atoi(idString)
	if err != nil {
		log.Println("Error converting topic ID to integer from request. urlVars:", urlVars)
		utils.ReturnHTTPErrorResponse(w, 305, "Error fetching topic")
		return
	}

	row := utils.DBCon.QueryRow(`
    SELECT
      id
			,name
    FROM main.topic
    WHERE id = $1
    `, topicID.ID)
	topic := new(models.Topic)
	if err := row.Scan(
		&topic.ID,
		&topic.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("Received request for nonexistent topic. TopicID:", topicID)
			utils.ReturnHTTPErrorResponse(w, 320, "Topic does not exist")
			return
		}
		log.Println("Error fetching requested topic from database. TopicID:", topicID, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 310, "Error fetching topic")
		return
	}

	if err := json.NewEncoder(w).Encode(topic); err != nil {
		log.Println("Error encoding response for requested topic. Topic:", topic, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 340, "Error fetching topic")
		return
	}
}

// NewTopic swagger:route POST /api/v1/topic Topics NewTopic
//
// Constructs and persists a new Topic based on the provided parameters
//
// Responses:
//    default: errorResponse
//        200: models.Topic
func NewTopic(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.NewTopicParam
	err := decoder.Decode(&t)
	if err != nil {
		log.Println("Error decoding request body for NewTopic. Body:", r.Body, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 200, "Error creating new topic")
		return
	}

	tx, err := utils.DBCon.Begin()
	if err != nil {
		log.Println("Error starting new transaction for inserting topic. Error:", err, "-- Topic:", t)
		utils.ReturnHTTPErrorResponse(w, 210, "Error creating new topic")
		return
	}
	defer tx.Rollback()

	var topicInsertID int
	err = tx.QueryRow(`
    INSERT INTO main.topic
    (name)
    VALUES ($1)
    returning id;
  `, t.Name).Scan(&topicInsertID)
	if err != nil {
		if err.(*pq.Error).Code.Name() == "unique_violation" {
			constraintPattern, _ := regexp.Compile("violates unique constraint \"(.+?)\"")
			matchSlice := constraintPattern.FindSubmatch([]byte(err.(*pq.Error).Message))
			if len(matchSlice) < 2 {
				log.Println("Received request to create topic but encountered unknown unique constraint violation. Error:", err)
				utils.ReturnHTTPErrorResponse(w, 230, "Unknown error attempting to create topic")
				return
			}
			constraint := string(matchSlice[1])
			if constraint == "topic_unq_name" {
				log.Println("Received request to create topic with duplicate name. Name:", t.Name)
				utils.ReturnHTTPErrorResponse(w, 240, "Topic name already in use")
				return
			}
			// TODO: Is there a better way to organize this?
		}
		log.Println("Error inserting new topic. Error:", err, "-- Topic:", t)
		utils.ReturnHTTPErrorResponse(w, 260, "Error creating new topic")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction for new topic. Error:", err, "-- Topic:", t)
		utils.ReturnHTTPErrorResponse(w, 280, "Error creating new topic")
		return
	}
	log.Println("Successfully created new topic")
}

// GetTopicsLike swagger:route GET /api/v1/topic_like/{part} Topics GetTopicsLike
//
// Gets topics similar to the given string
//
// Responses:
//    default: errorResponse
//        200: models.Topic
func GetTopicsLike(w http.ResponseWriter, r *http.Request) {
	urlVars := mux.Vars(r)
	partString, ok := urlVars["part"]
	if !ok {
		log.Println("Error retrieving topic part from request. urlVars:", urlVars)
		utils.ReturnHTTPErrorResponse(w, 300, "Error retrieving similar topic")
		return
	}

	topicSlice := make([]*models.Topic, 0)
	rows, err := utils.DBCon.Query(`
    SELECT
      id
			,name
    FROM main.topic
    WHERE name LIKE '%' || $1 || '%'
    `, partString)
	if err != nil {
		log.Println("Error retrieving similar topics from database. TopicPart:", partString, " -- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 320, "Error retrieving similar topic")
		return
	}
	for rows.Next() {
		topic := new(models.Topic)
		if err := rows.Scan(
			&topic.ID,
			&topic.Name,
		); err != nil {
			log.Println("Error scanning similar topic from database. TopicPart:", partString, "-- Error:", err)
			utils.ReturnHTTPErrorResponse(w, 310, "Error retrieving similar topic")
			continue
		}
		topicSlice = append(topicSlice, topic)
	}

	if err := json.NewEncoder(w).Encode(topicSlice); err != nil {
		log.Println("Error encoding response for requested topics. TopicSlice:", topicSlice, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 340, "Error retrieving similar topic")
		return
	}

	// TODO: Do fancy fuzzy-matching stuff
}
