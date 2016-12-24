package models

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	Password string `json:"password,omitempty"`

	// the time the user was created
	Created time.Time `json:"created"`

	// the last time the user interacted with the site
	LastActivity time.Time `json:"last_activity"`

	// a brief self-description of the user
	Description string `json:"description"`

	// a URL pointing to the user's icon
	IconURL string `json:"icon_url"`
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
	Password string `json:"password,omitempty"`

	// a brief self-description of the user
	//
	// in: body
	Description string `json:"description"`

	// a URL pointing to the user's icon
	//
	// in: body
	IconURL string `json:"icon_url"`
}

// MarshalJSON allows us to Marshal Users such that the Created and
// LastActivity Times are represented in RFC 3339 format.
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Created      string `json:"created"`
		LastActivity string `json:"last_activity"`
		//Password     string `json:"password,omitempty"`
		*Alias
	}{
		Created:      fmt.Sprintf(`%s`, u.Created.Format(time.RFC3339)),
		LastActivity: fmt.Sprintf(`%s`, u.LastActivity.Format(time.RFC3339)),
		//Password:     "",
		Alias: (*Alias)(u),
	})
}

// UnmarshalJSON allows us to UnmarshalJSON Users such that the Created and
// LastActivity Times can be parsed from RFC 3339 time strings
func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		Created      string `json:"created"`
		LastActivity string `json:"last_activity"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.Created, _ = time.Parse(time.RFC3339, aux.Created)
	u.LastActivity, _ = time.Parse(time.RFC3339, aux.LastActivity)
	return nil
}
