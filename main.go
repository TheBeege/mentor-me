package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/TheBeege/mentor-me/models"

	"github.com/gorilla/mux"
)

var (
	dbCon *sql.DB
)

type errorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func main() {
	log.Println("It's go time!")
	dbHost := os.Getenv("dbhost")
	dbUser := os.Getenv("dbuser")
	dbPass := os.Getenv("dbpass")
	dbDatabase := os.Getenv("dbdatabase")
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbDatabase)

	var err error
	dbCon, err = sql.Open("postgres", conString)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return
	}
	defer dbCon.Close()

	router := mux.NewRouter().StrictSlash(true)
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./client/"))))
	router.HandleFunc("/", Index).Methods("GET")

	router.HandleFunc("/api/v1/user/{id:[0-9]+}", GetUser).Methods("GET")
	router.HandleFunc("/api/v1/user", NewUser).Methods("POST")

	router.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))

	log.Println("Time to serve it up!")
	log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("kthxbye")
}

// Index should return hello world
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// GetUser swagger:route GET /api/v1/user/{id} user getUser
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
		log.Println("Error retrieving user ID from request. urlVars: ", urlVars)
		returnHTTPErrorResponse(w, 300, "Error fetching user")
		return
	}

	var err error
	userID.ID, err = strconv.Atoi(idString)
	if err != nil {
		log.Println("Error converting user ID to integer from request. urlVars: ", urlVars)
		returnHTTPErrorResponse(w, 305, "Error fetching user")
		return
	}

	row := dbCon.QueryRow(`
    SELECT
      user_id
      ,username
      ,display_name
      ,email
      ,created
    FROM main.user
    WHERE user_id = $1
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
			log.Println("Received request for nonexistent user. UserID: ", userID)
			returnHTTPErrorResponse(w, 320, "User does not exist")
			return
		} else {
			log.Println("Error fetching requested user from database. UserID: ", userID, " -- Error: ", err)
			returnHTTPErrorResponse(w, 310, "Error fetching user")
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Error encoding response for requested user. User: ", user, " -- Error: ", err)
		returnHTTPErrorResponse(w, 340, "Error fetching user")
		return
	}
}

// NewUser constructs and persists a new Channel based on the provided parameters
// swagger:route POST /api/v1/user user newUser
//
// Creates a new user.
//
// Responses:
//    default: errorResponse
//        200: models.User
func NewUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var u models.NewUserParam
	err := decoder.Decode(&u)
	if err != nil {
		log.Println("Error decoding request body for NewUser. Body: ", r.Body, " -- Error: ", err)
		returnHTTPErrorResponse(w, 200, "Error creating new user")
		return
	}
	// TODO: Encrypt user password

	tx, err := dbCon.Begin()
	if err != nil {
		log.Println("Error starting new transaction for inserting new channel. Error: ", err, " -- User: ", u)
		returnHTTPErrorResponse(w, 210, "Error creating new user")
		return
	}
	defer tx.Rollback()

	var userInsertID int
	err = tx.QueryRow(`
    INSERT INTO main.user
    (username, display_name, email, password, created, last_activity)
    VALUES ($1, $2, $3, $4, now(), now())
    returning user_id;
  `, u.Username, u.DisplayName, u.Email, u.Password).Scan(&userInsertID)
	// TODO: logic for unique name/email collisions
	if err != nil {
		log.Println("Error inserting new user. Error: ", err, " -- User: ", u)
		returnHTTPErrorResponse(w, 220, "Error creating new user")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("Error committing transaction for new user. Error: ", err, " -- User: ", u)
		returnHTTPErrorResponse(w, 230, "Error creating new user")
		return
	}
	log.Println("Successfully created new user")
}

func returnHTTPErrorResponse(w http.ResponseWriter, code int, message string) {
	r := &errorResponse{ErrorCode: code, ErrorMessage: message}
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("Failed to write HTTP Error Response. Error: ", err)
	}
}
