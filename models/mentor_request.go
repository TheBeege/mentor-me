package models

import (
	"encoding/json"
	"fmt"
	"time"
)

// MentorRequest represents a mentee's request to enter a mentoring relationship.
// swagger:model
type MentorRequest struct {
	// the mentor user
	//
	// required: true
	// read only: true
	Mentor *User `json:"mentor"`

	// the mentee user
	//
	// required: true
	// read only: true
	Mentee *User `json:"mentee"`

	// the timestamp of when this request was created
	//
	// read only: true
	Requested time.Time `json:"requested"`

	// the timestamp of when this request was accepted, if it has been accepted
	Accepted time.Time `json:"accepted"`

	// the timestamp of when this request was rejected, if it has been rejected
	Rejected time.Time `json:"rejected,omitempty"`
}

// NewMentorRequestParam is used for the NewMentorRequest API operation
// swagger:parameters NewMentorRequest
type NewMentorRequestParam struct {
	// the ID of the mentor user
	//
	// in: body
	MentorID int `json:"mentor_id"`

	// the ID of the mentee user
	//
	// unique: true
	// in: body
	MenteeID int `json:"mentee_id"`
}

// MarshalJSON allows us to Marshal Users such that the Created and
// LastActivity Times are represented in RFC 3339 format.
func (mr *MentorRequest) MarshalJSON() ([]byte, error) {
	type Alias MentorRequest
	return json.Marshal(&struct {
		Requested string `json:"requested"`
		Accepted  string `json:"accepted"`
		Rejected  string `json:"rejected"`
		*Alias
	}{
		Requested: fmt.Sprintf(`%s`, mr.Requested.Format(time.RFC3339)),
		Accepted:  fmt.Sprintf(`%s`, mr.Accepted.Format(time.RFC3339)),
		Rejected:  fmt.Sprintf(`%s`, mr.Rejected.Format(time.RFC3339)),
		Alias:     (*Alias)(mr),
	})
}

// UnmarshalJSON allows us to UnmarshalJSON Users such that the Created and
// LastActivity Times can be parsed from RFC 3339 time strings
func (mr *MentorRequest) UnmarshalJSON(data []byte) error {
	type Alias MentorRequest
	aux := &struct {
		Requested string `json:"requested"`
		Accepted  string `json:"accepted"`
		Rejected  string `json:"rejected"`
		*Alias
	}{
		Alias: (*Alias)(mr),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	mr.Requested, _ = time.Parse(time.RFC3339, aux.Requested)
	mr.Accepted, _ = time.Parse(time.RFC3339, aux.Accepted)
	mr.Rejected, _ = time.Parse(time.RFC3339, aux.Rejected)
	return nil
}
