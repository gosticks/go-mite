package mite_test

import (
	"os"
	"testing"

	mite "github.com/gosticks/go-mite"
)

func TestGetProjects(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	_, err := mite.GetProjects(nil)
	if err != nil {
		t.Error(username, team, key, err)
	}
}

func TestGetProject(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	ps, err := mite.GetProjects(nil)
	if err != nil {
		t.Error(username, team, key, err)
	}

	for _, p := range ps {
		_, errUser := mite.GetProject(p.ID)
		if errUser != nil {
			t.Errorf("Failed to get the service for entry (name=%s id=%d)", p.Name, p.ID)
			t.FailNow()
		}
	}

	t.Logf("Got all projects available for account by id (total %d)", len(ps))
}
