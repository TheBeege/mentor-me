package models

import (
	"fmt"
	"strings"
	"time"
)

// JSONTime is a wrapper around time.Time to allow JSON unmarshalling of time
// values in RFC3339 format  http://stackoverflow.com/a/23695774
type JSONTime struct {
	time.Time
}

// MarshalJSON allows JSONTime to fulfill the Marshaler interface for
// marshalling JSONTime values into JSON strings
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf(`"%s"`, t.Time.Format(time.RFC3339))
	return []byte(stamp), nil
}

// UnmarshalJSON allows JSONTime to fulfill the Unmarshaler interface for
// unmarshalling JSON strings (byte slices) into time.Time (JSONTime) values
func (t *JSONTime) UnmarshalJSON(jsonString []byte) error {
	timeVal, err := time.Parse(time.RFC3339, strings.Trim(string(jsonString), `"`))
	if err != nil {
		return err
	}
	t.Time = timeVal
	return nil
}
