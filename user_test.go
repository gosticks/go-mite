package mite_test

import (
	"os"
	"testing"

	"github.com/gosticks/go-mite"
)

func TestGetUsers(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	_, errUser := mite.GetUsers()
	if errUser != nil {
		t.Error(username, team, key, errUser)
	}
}
