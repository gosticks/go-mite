package mite_test

import (
	"os"
	"testing"

	mite "github.com/gosticks/go-mite"
)

func TestGetServices(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	_, err := mite.GetServices(nil)
	if err != nil {
		t.Error(username, team, key, err)
	}
}

func TestGetService(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	mite := mite.NewMiteAPI(username, team, key, "test@go-mite")

	services, err := mite.GetServices(nil)
	if err != nil {
		t.Error(username, team, key, err)
	}

	for _, s := range services {
		_, errUser := mite.GetService(s.ID)
		if errUser != nil {
			t.Errorf("Failed to get the service for entry (name=%s id=%d)", s.Name, s.ID)
			t.FailNow()
		}
	}

	t.Logf("Got all services available for account by id (total %d)", len(services))
}
