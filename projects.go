package mite

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// -------------------------------------------------------------
// ~ API Mappings
// -------------------------------------------------------------

// Project mite object
type Project struct {
	ID                    uint64               `json:"id"`
	Name                  string               `json:"name"`
	Note                  string               `json:"note"`
	CustomerID            uint64               `json:"customer_id"`
	CustomerName          string               `json:"customer_name"`
	Budget                uint64               `json:"budget"`
	BudgetType            string               `json:"budgetType"`
	HourlyRate            uint64               `json:"hourly_rate"`
	Archived              bool                 `json:"archived"`
	ArchivedHourlyRata    string               `json:"active_hourly_rate"`
	HourlyRatesPerService []ServiceHourlyRates `json:"hourly_rates_per_service"`
	CreatedAt             time.Time            `json:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at"`
}

func (p *Project) String() string {
	return fmt.Sprintf("%d: %s for %s (archived: %t)", p.ID, p.Name, p.CustomerName, p.Archived)
}

type getProjectsResponseWrapper struct {
	Project *Project `json:"project"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

// GetProjects returns mite projects
// possible filters https://mite.yo.lk/en/api/projects.html
func (m *Mite) GetProjects(filters map[string]string) ([]*Project, error) {
	var projectResponse []*getProjectsResponseWrapper
	err := m.getAndDecodeFromSuffix("projects.json", &projectResponse, filters)
	if err != nil {
		return nil, err
	}

	projects := make([]*Project, len(projectResponse))

	// Unwrap all the data
	for i, p := range projectResponse {
		projects[i] = p.Project
		// fmt.Println("Project", p.Project)
	}

	return projects, nil
}

// GetProject returns a single project
func (m *Mite) GetProject(id uint64) (*Project, error) {
	var resp *getProjectsResponseWrapper
	err := m.getAndDecodeFromSuffix("projects/"+strconv.FormatUint(id, 10)+".json", &resp, nil)
	if err != nil {
		return nil, err
	}
	return resp.Project, nil
}

// -------------------------------------------------------------
// ~ Create
// -------------------------------------------------------------

// CreateProject creates a project
func (m *Mite) CreateProject(project *Project) (*Project, error) {
	reqData := &getProjectsResponseWrapper{Project: project}
	resp, errRequest := m.postToMite("/projects.json", reqData)
	if errRequest != nil {
		return nil, errRequest
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Failed to create a time entry: " + resp.Status)
	}

	var respEntry = &getProjectsResponseWrapper{}

	// Unmarshal data
	err := json.NewDecoder(resp.Body).Decode(respEntry)
	if err != nil {
		return nil, err
	}

	return respEntry.Project, nil
}

// DeleteProject deletes a project from mite if the project has time entries it cannot be deleted
// please archive it instead or delete all the entries first
func (m *Mite) DeleteProject(id uint64) error {
	res, errReq := m.deleteFromMite("/projects/"+strconv.FormatUint(id, 10)+".json", nil)
	if errReq != nil {
		return errReq
	}

	if res.StatusCode == http.StatusOK {
		return nil
	}

	if res.StatusCode == http.StatusUnprocessableEntity {
		return fmt.Errorf("failed to delete project since it containes time entries please archive")
	}

	return fmt.Errorf("unknown delete error code %d", res.StatusCode)
}
