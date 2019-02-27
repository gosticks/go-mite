package mite_test

import (
	"os"
	"testing"
	"time"

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

	// should fail
	_, errEntries := m.GetTimeEntriesGroup(now.BeginningOfMonth(), now.EndOfMonth(), nil)
	if errEntries == nil {
		t.Error(errEntries)
	}

	// should not fail
	entries, errEntries := m.GetTimeEntriesGroup(now.BeginningOfMonth(), now.EndOfMonth(), map[string]string{"group_by": "user"})
	if errEntries != nil {
		t.Error(errEntries)
	}

	t.Logf("loaded %d groups", len(entries))
}

func TestCreateDeleteTimeEntry(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	m := mite.NewMiteAPI(username, team, key, "test@go-mite")

	newEntry := &mite.TimeEntryCreator{
		DateAt:  mite.Time{time.Now()},
		Minutes: 60,
		Note:    "TEST NOTE CREATED BY GO-MITE API. PLEASE REMOVE",
	}

	entry, errEntry := m.CreateTimeEntry(newEntry)
	if errEntry != nil {
		t.Error(errEntry)
	}

	t.Logf("Created entry: %d", entry.ID)

	errDel := m.DeleteTimeEntry(entry.ID)
	if errDel != nil {
		t.Error(errDel)
	}
}
