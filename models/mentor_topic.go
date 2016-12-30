package models

// MentorTopic represents a user's interest in mentoring in the given topic.
// swagger:model
type MentorTopic struct {
	// the ID of the mentor/user
	//
	// required: true
	// read only: true
	UserID int `json:"user_id"`

	// the ID of the topic
	//
	// required: true
	// read only: true
	TopicID int `json:"topic_id"`

	// the self-assessed level of the mentor's skill in the topic
	//
	// required: true
	// pattern: [1-5]
	Level int `json:"level"`

	// a brief description of the mentor's experience in the topic
	//
	// required: true
	Description string `json:"description"`
}

// NewMentorTopicParam is used in the NewMentorTopic API operation
// swagger:parameters NewMentorTopic
type NewMentorTopicParam struct {
	// the ID of the mentor/user
	//
	// required: true
	// in: body
	UserID int `json:"user_id"`

	// the ID of the topic
	//
	// required: true
	// in: body
	TopicID int `json:"topic_id"`

	// the self-assessed level of the mentor's skill in the topic
	//
	// pattern: [1-5]
	// required: true
	// in: body
	Level int `json:"level"`

	// a brief description of the mentor's experience in the topic
	//
	// required: true
	// in: body
	Description string `json:"description"`
}

// FindMentorsParam is used in the FindMentors API operation
// swagger:parameters FindMentors
type FindMentorsParam struct {
	// the name of the desired topic area
	//
	// in: path
	// required: true
	TopicName string `json:"topic"`

	// the desired level of proficiency in the desired topic
	//
	// in: path
	// pattern: [1-5]
	// required: true
	Level int `json:"level"`
}
