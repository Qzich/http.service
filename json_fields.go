package http_service

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONInt struct {
	Value int
	Valid bool
	Set   bool
}

func (i *JSONInt) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	i.Set = true

	if string(data) == "null" {
		// The key was set to null
		i.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	i.Valid = true
	return nil
}

type JSONString struct {
	Value string
	Valid bool
	Set   bool
}

func (str *JSONString) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	str.Set = true

	if string(data) == "null" {
		// The key was set to null
		str.Valid = false
		return nil
	}

	// The key isn't set to null
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	str.Value = temp
	str.Valid = true
	return nil
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}
