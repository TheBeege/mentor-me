package models

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshalJSONTimeHappyPath(t *testing.T) {
	type TestTimeStruct struct {
		TestTime JSONTime `json:"time"`
	}
	nowTime := time.Now()
	testStruct := TestTimeStruct{
		TestTime: JSONTime{Time: nowTime},
	}
	byteSlice, err := json.Marshal(testStruct)
	if err != nil {
		t.Error("Failed to marshal JSONTime to string")
		return
	}
	expectedString := fmt.Sprintf("{\"time\":\"%s\"}", nowTime.Format(time.RFC3339))
	if string(byteSlice) != expectedString {
		t.Errorf(`Marshalling JSONTime did not output correct format. Expected "%s". Output was "%s".`, expectedString, string(byteSlice))
		return
	}
}

func TestUnmarshalJSONTimeHappyPath(t *testing.T) {
	type TestTimeStruct struct {
		TestTime JSONTime `json:"time"`
	}
	var testStruct TestTimeStruct
	testByteArray := []byte(`{"time":"2016-02-20T20:36:21-08:00"}`)
	err := json.Unmarshal(testByteArray, &testStruct)
	if err != nil {
		t.Errorf("Failed to unmarshal test struct from JSONTime: %s", err)
		return
	}
	if testStruct.TestTime.Format(time.RFC3339) != "2016-02-20T20:36:21-08:00" {
		t.Errorf(`Unmarshaling text "%s" did not yield a proper time. Found "%s"`,
			string(testByteArray), testStruct.TestTime.Format(time.RFC3339))
	}
}

func TestUnmarshalJSONTimeBadFormat(t *testing.T) {
	type TestTimeStruct struct {
		TestTime JSONTime `json:"time"`
	}
	var testStruct TestTimeStruct
	testByteArray := []byte(`{"time":"2016-02-20T20:36:21-08:000000000"}`)
	err := json.Unmarshal(testByteArray, &testStruct)
	if err == nil {
		t.Error("Unmarshalled a badly formatted time.")
		return
	}
}
