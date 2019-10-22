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

	_, errUser := mite.GetUsers(false)
	if errUser != nil {
		t.Error(username, team, key, errUser)
	}

	_, errArchivedUser := mite.GetUsers(true)
	if errArchivedUser != nil {
		t.Error(username, team, key, errUser)
	}
}
func TestGetUser(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	users, errUser := mite.GetUsers(false)
	if errUser != nil {
		t.Error(username, team, key, errUser)
	}

	for _, user := range users {
		_, errUser := mite.GetUser(user.ID)
		if errUser != nil {
			t.Errorf("Failed to get the user available in the user map username=%s id=%d", user.Name, user.ID)
			t.FailNow()
		}
	}

	t.Logf("Got all users available for account by id (total %d)", len(users))
}
