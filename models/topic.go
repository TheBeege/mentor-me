package models

/*import (
	"encoding/json"
	"fmt"
	"time"
)*/

// Topic represents an an area of study between mentors and mentees.
// swagger:model
type Topic struct {
	// the ID of the topic
	//
	// required: true
	// read only: true
	ID int `json:"id"`

	// the name of the topic
	//
	// required: true
	Name string `json:"name"`
}

// TopicIDParam is used for the GetTopic API operation
// swagger:parameters getTopic
type TopicIDParam struct {
	// the ID of the topic
	//
	// unique: true
	// in: path
	ID int `json:"id"`
}

// NewTopicParam is used for the NewTopic API operation
// swagger:parameters newTopic
type NewTopicParam struct {
	// the name of the new topic
	//
	// unique: true
	// in: body
	Name string `json:"name"`
}

// TopicLikeParam is used for the GetTopicLike API operation
// swagger:parameters getTopicsLike
type TopicLikeParam struct {
	// string to search for similar topics
	//
	// in:path
	Part string `json:"part"`
}
