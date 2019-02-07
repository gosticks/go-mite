package mite

import (
	"fmt"
	"strings"
	"time"
)

// Time is the mite time object
type Time struct {
	time.Time
}

// UnmarshalYAML unmarshals mite time from yaml
func (mt *Time) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var s string
	if err = unmarshal(&s); err != nil {
		return
	}

	mt.Time, err = time.Parse(TimeFormat, s)
	return
}

// UnmarshalJSON unmarshals mite time from json
func (mt *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		mt.Time = time.Time{}
		return
	}
	mt.Time, err = time.Parse(TimeFormat, s)
	return
}

// MarshalJSON marshals mite time to json
func (mt *Time) MarshalJSON() ([]byte, error) {
	if mt.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", mt.Time.Format(TimeFormat))), nil
}

// TimeString returns a formated string for a mite time
func TimeString(t time.Time) string {
	return fmt.Sprintf("%s", t.Format(TimeFormat))
}

func (mt *Time) String() string {
	return mt.Time.Format(TimeFormat)
}
