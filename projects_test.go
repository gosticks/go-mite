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

func TestCreateDeleteProject(t *testing.T) {
	username, okUser := os.LookupEnv("MITE_USER")
	team, okAddr := os.LookupEnv("MITE_TEAM")
	key, okKey := os.LookupEnv("MITE_APIKEY")
	if !okAddr || !okUser || !okKey {
		t.Errorf("username=%s, team=%s and key=%s are required", username, team, key)
		t.FailNow()
	}

	m := mite.NewMiteAPI(username, team, key, "test@go-mite")

	cp := &mite.Project{
		Name: "GoTestProject",
	}

	p, errCreate := m.CreateProject(cp)
	if errCreate != nil {
		t.Error("failed to create project", errCreate)
		t.FailNow()
	}

	t.Logf("Project created ID=%d", p.ID)

	// Deleting project
	errDelete := m.DeleteProject(p.ID)
	if errDelete != nil {
		t.Errorf("failed to delete project please clean it up for me... sorry ID=%d Name=%s", p.ID, p.Name)
		t.FailNow()
	}
}
