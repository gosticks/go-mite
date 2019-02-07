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

// TimeEntryGroup mapping to the mite return type
type TimeEntryGroup struct {
	Minutes  uint64                 `json:"minutes"`
	Revenue  float64                `json:"revenue"`
	UserID   uint64                 `json:"user_id"`
	UserName string                 `json:"user_name"`
	From     Time                   `json:"from"`
	To       Time                   `json:"to"`
	Params   map[string]interface{} `json:"time_entries_params"`
}

// TimeEntry mapping to the mite return type
type TimeEntry struct {
	ID       uint64 `json:"id"`
	Minutes  uint64 `json:"minutes"`
	DateAt   Time   `json:"date_at"`
	Note     string `json:"note"`
	Billable bool   `json:"billable"`
	Locked   bool   `json:"locked"`
	// Revenue bool `json:"locked"`
	HourlyRate   uint64    `json:"hourly_rate"`
	UserID       uint64    `json:"user_id"`
	UserName     string    `json:"user_name"`
	ProjectID    uint64    `json:"project_id"`
	ProjectName  string    `json:"project_name"`
	CustomerID   uint64    `json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	ServiceID    uint64    `json:"service_id"`
	ServiceName  string    `json:"service_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type getTimeEntryResponseWrapper struct {
	TimeEntry *TimeEntry `json:"time_entry"`
}

type getTimeEntriesGroupResponseWrapper struct {
	TimeEntry *TimeEntryGroup `json:"time_entry_group"`
}

// -------------------------------------------------------------
// ~ Get
// -------------------------------------------------------------

// GetTimeEntriesGroup return time entry groups for a timerange with filters
func (m *Mite) GetTimeEntriesGroup(from, to time.Time, filters map[string]string) ([]*TimeEntryGroup, error) {
	var timeGroupEntries []*getTimeEntriesGroupResponseWrapper
	err := m.getAndDecodeFromSuffix("time_entries.json", &timeGroupEntries, filters)
	if err != nil {
		return nil, err
	}
	timeEntries := make([]*TimeEntryGroup, len(timeGroupEntries))

	// Unwrap all the data
	for i, t := range timeGroupEntries {
		timeEntries[i] = t.TimeEntry
		// fmt.Println("Time Entry", t.TimeEntry)
	}
	return timeEntries, nil
}

// GetTimeEntries returns arrays for a time range
func (m *Mite) GetTimeEntries(from, to time.Time, filters map[string]string) ([]*TimeEntry, error) {
	var timeResp []*getTimeEntryResponseWrapper
	err := m.getAndDecodeFromSuffix("time_entries.json", &timeResp, filters)
	if err != nil {
		return nil, err
	}
	timeEntries := make([]*TimeEntry, len(timeResp))

	// Unwrap all the data
	for i, t := range timeResp {
		timeEntries[i] = t.TimeEntry
		// fmt.Println("Time Entry", t.TimeEntry)
	}
	return timeEntries, nil
}

// GetTimeEntriesForProjectByService returns a array of time entry groups grouped by service
func (m *Mite) GetTimeEntriesForProjectByService(from, to time.Time, projectID uint64) ([]*TimeEntryGroup, error) {
	params := map[string]string{
		ParamProjectID: fmt.Sprint(projectID),
		ParamGroupBy:   "service",
	}

	return m.GetTimeEntriesGroup(from, to, params)
}

// GetTimeEntry returns a time entry for a id
func (m *Mite) GetTimeEntry(id uint64) (*TimeEntry, error) {
	var timeResp *getTimeEntryResponseWrapper
	err := m.getAndDecodeFromSuffix("time_entries/"+strconv.FormatUint(id, 10)+".json", &timeResp, nil)
	if err != nil {
		return nil, err
	}
	return timeResp.TimeEntry, nil
}

// -------------------------------------------------------------
// ~ Create
// -------------------------------------------------------------

// CreateTimeEntry  creates a new time entry
func (m *Mite) CreateTimeEntry(entry *TimeEntry) (*TimeEntry, error) {
	reqData := &getTimeEntryResponseWrapper{TimeEntry: entry}
	resp, errRequest := m.postToMite("/time_entries.json", reqData)
	if errRequest != nil {
		return nil, errRequest
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Failed to create a time entry: " + resp.Status)
	}

	var respEntry = &getTimeEntryResponseWrapper{}

	// Unmarshal data
	err := json.NewDecoder(resp.Body).Decode(respEntry)
	if err != nil {
		return nil, err
	}

	return respEntry.TimeEntry, nil
}

// -------------------------------------------------------------
// ~ Update
// -------------------------------------------------------------

// UpdateTimeEntry updates the fields provided in the update struct for a id
func (m *Mite) UpdateTimeEntry(id uint64, update *TimeEntry) error {
	// Wrap time entry
	wrap := &getTimeEntryResponseWrapper{TimeEntry: update}

	resp, errRequest := m.patchAtMite("/time_entries/"+strconv.FormatUint(id, 10)+".json", wrap)
	if errRequest != nil {
		return errRequest
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to create a time entry: " + resp.Status)
	}
	return nil
}

// -------------------------------------------------------------
// ~ Delete
// -------------------------------------------------------------

// DeleteTimeEntry removes a time entry for a user
func (m *Mite) DeleteTimeEntry(id uint64) error {
	resp, errRequest := m.deleteFromMite("/time_entries/"+strconv.FormatUint(id, 10)+".json", nil)
	if errRequest != nil {
		return errRequest
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to create a time entry: " + resp.Status)
	}
	return nil
}
