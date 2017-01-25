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

// GetUser swagger:route GET /api/v1/user/{id} Users GetUser
//
// Gets info for a user.
//
// Responses:
//    default: errorResponse
//        200: models.User
func GetUser(w http.ResponseWriter, r *http.Request) {
	var userID models.UserIDParam
	urlVars := mux.Vars(r)
	idString, ok := urlVars["id"]
	if !ok {
		log.Println("Error retrieving user ID from request. urlVars:", urlVars)
		utils.ReturnHTTPErrorResponse(w, 300, "Error fetching user")
		return
	}

	var err error
	userID.ID, err = strconv.Atoi(idString)
	if err != nil {
		log.Println("Error converting user ID to integer from request. urlVars:", urlVars)
		utils.ReturnHTTPErrorResponse(w, 305, "Error fetching user")
		return
	}

	row := utils.DBCon.QueryRow(`
    SELECT
      id
      ,username
      ,display_name
      ,email
      ,created
    FROM main.user
    WHERE id = $1
    `, userID.ID)
	user := new(models.User)
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.DisplayName,
		&user.Email,
		&user.Created,
	); err != nil {
		if err == sql.ErrNoRows {
			log.Println("Received request for nonexistent user. UserID:", userID)
			utils.ReturnHTTPErrorResponse(w, 320, "User does not exist")
			return
		}
		log.Println("Error fetching requested user from database. UserID:", userID, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 310, "Error fetching user")
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error encoding response for requested user. User:", user, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 340, "Error fetching user")
		return
	}
}

// NewUser swagger:route POST /api/v1/user Users NewUser
//
// Constructs and persists a new user based on the provided parameters
//
// Responses:
//    default: errorResponse
//        200: models.User
func NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u models.NewUserParam
	err := decoder.Decode(&u)
	if err != nil {
		log.Println("Error decoding request body for NewUser. Body:", r.Body, "-- Error:", err)
		utils.ReturnHTTPErrorResponse(w, 200, "Error creating new user")
		return
	}
	// TODO: Encrypt user password

	tx, err := utils.DBCon.Begin()
	if err != nil {
		log.Println("Error starting new transaction for inserting new channel. Error:", err, "-- User:", u)
		utils.ReturnHTTPErrorResponse(w, 210, "Error creating new user")
		return
	}
	defer tx.Rollback()

	var userInsertID int
	err = tx.QueryRow(`
    INSERT INTO main.user
    (username, display_name, email, password, created, last_activity)
    VALUES ($1, $2, $3, $4, now(), now())
    returning id;
  `, u.Username, u.DisplayName, u.Email, u.Password).Scan(&userInsertID)
	if err != nil {
		if err.(*pq.Error).Code.Name() == "unique_violation" {
			constraintPattern, _ := regexp.Compile("violates unique constraint \"(.+?)\"")
			matchSlice := constraintPattern.FindSubmatch([]byte(err.(*pq.Error).Message))
			if len(matchSlice) < 2 {
				log.Println("Received request to create user but encountered unknown unique constraint violation. Error:", err)
				utils.ReturnHTTPErrorResponse(w, 230, "Unknown error attempting to create user")
				return
			}
			constraint := string(matchSlice[1])
			if constraint == "user_unq_username" {
				log.Println("Received request to create user with duplicate username. Username:", u.Username)
				utils.ReturnHTTPErrorResponse(w, 240, "Username already in use")
				return
			} else if constraint == "user_unq_email" {
				log.Println("Received request to create user with duplicate email. Email:", u.Email)
				utils.ReturnHTTPErrorResponse(w, 250, "Email already in use")
				return
			}
		}
		log.Println("Error inserting new user. Error:", err, "-- User:", u)
		utils.ReturnHTTPErrorResponse(w, 260, "Error creating new user")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction for new user. Error:", err, "-- User:", u)
		utils.ReturnHTTPErrorResponse(w, 280, "Error creating new user")
		return
	}
	log.Println("Successfully created new user")
}
