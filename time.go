package mite

import (
	"fmt"
	"strings"
	"time"
)

type MiteTime struct {
	time.Time
}

func (mt *MiteTime) UnmarshalYAML(unmarshal func(interface{}) error) (err error) {
	var s string
	if err = unmarshal(&s); err != nil {
		return
	}

	mt.Time, err = time.Parse(MiteTimeFormat, s)
	return
}

func (mt *MiteTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		mt.Time = time.Time{}
		return
	}
	mt.Time, err = time.Parse(MiteTimeFormat, s)
	return
}

func (mt *MiteTime) MarshalJSON() ([]byte, error) {
	if mt.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", mt.Time.Format(MiteTimeFormat))), nil
}

func MiteTimeString(t time.Time) string {
	return fmt.Sprintf("%s", t.Format(MiteTimeFormat))
}

func (mt *MiteTime) String() string {
	return mt.Time.Format(MiteTimeFormat)
}
