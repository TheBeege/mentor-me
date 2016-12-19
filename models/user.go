package models

// User represents an end user on the site. Includes all info necessary to identify a user and includes that user's playlists.
// swagger:model
type User struct {
	// the ID of the user
	//
	// required: true
	// read only: true
	ID int `json:"id"`

	// the username of the user
	//
	// required: true
	Username string `json:"username"`

	// the way a user's name will be displayed, a nickname, if you will
	DisplayName string `json:"display_name"`

	// the user's email
	//
	// required: true
	// pattern: [^@]+@[\w\d\.]+
	Email string `json:"email"`

	// the user's password. we may be able to avoid doing this. TODO: check
	Password string `json:"password"`

	// the time the user was created
	Created JSONTime `json:"created"`

	// the last time the user interacted with the site
	LastActivity JSONTime `json:"last_activity"`
}

// UserIDParam is used for the GetUser API operation
// swagger:parameters getUser
type UserIDParam struct {
	// the ID of the user
	//
	// unique: true
	// in: path
	ID int `json:"id"`
}

// NewUserParam is used for the NewUser API operation
// swagger:parameters newUser
type NewUserParam struct {
	// the username of the new user
	//
	// unique: true
	// in: body
	Username string `json:"username"`

	// the pretty name or nickname of the new user
	//
	// in: body
	DisplayName string `json:"display_name"`

	// the email of the new user
	//
	// unique: true
	// in: body
	Email string `json:"email"`

	// the password of the new user
	//
	// in: body
	Password string `json:"password"`
}
