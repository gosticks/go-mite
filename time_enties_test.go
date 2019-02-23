package mite_test

import (
	"os"
	"testing"

	"github.com/jinzhu/now"

	mite "github.com/gosticks/go-mite"
)

func TestGetTimeEntriesGroup(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	m := mite.NewMiteAPI(username, team, key, "test@go-mite")

	entries, errEntries := m.GetTimeEntriesGroup(now.BeginningOfMonth(), now.EndOfMonth(), map[string]string{"group_by": "user"})
	if errEntries != nil {
		t.Error(errEntries)
	}

	t.Logf("loaded %d groups", len(entries))
}
