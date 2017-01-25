package main

import (
	"database/sql"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"github.com/TheBeege/mentor-me/controllers"
	"github.com/TheBeege/mentor-me/utils"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("It's go time!")
	dbHost := os.Getenv("dbhost")
	dbUser := os.Getenv("dbuser")
	dbPass := os.Getenv("dbpass")
	dbDatabase := os.Getenv("dbdatabase")
	conString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbDatabase)

	var err error
	utils.DBCon, err = sql.Open("postgres", conString)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return
	}
	defer utils.DBCon.Close()

	router := mux.NewRouter().StrictSlash(true)
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./client/"))))
	router.HandleFunc("/", Index).Methods("GET")

	router.HandleFunc("/api/v1/user/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/user", controllers.NewUser).Methods("POST")

	router.HandleFunc("/api/v1/topic/{id:[0-9]+}", controllers.GetTopic).Methods("GET")
	router.HandleFunc("/api/v1/topic", controllers.NewTopic).Methods("POST")
	router.HandleFunc("/api/v1/topic_like/{part:.+}", controllers.GetTopicsLike).Methods("GET")

	router.HandleFunc("/api/v1/mentor_topic", controllers.NewMentorTopic).Methods("POST")
	router.HandleFunc("/api/v1/find_mentors/{topic:.+}/{level:[1-5]}", controllers.FindMentors).Methods("GET")

	router.HandleFunc("/api/v1/mentor_request", controllers.NewMentorRequest).Methods("POST")
	router.HandleFunc("/api/v1/mentor_request/mentor/{mentor_id:[0-9]}", controllers.GetMentorRequestList).Methods("GET")
	router.HandleFunc("/api/v1/mentor_request/{id:[0-9]}", controllers.GetMentorRequest).Methods("GET")
	router.HandleFunc("/api/v1/mentor_request/{id:[0-9]}/accept", controllers.AcceptMentorRequest).Methods("POST")
	router.HandleFunc("/api/v1/mentor_request/{id:[0-9]}/reject", controllers.RejectMentorRequest).Methods("POST")

	router.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))

	log.Println("Time to serve it up!")
	log.Fatal(http.ListenAndServe(":8080", router))
	log.Println("kthxbye")
}

// Index should return hello world
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
