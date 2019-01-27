package mite

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// -------------------------------------------------------------
// ~ API Mappings
// -------------------------------------------------------------

type TimeEntryGroup struct {
	Minutes  uint64                 `json:"minutes"`
	Revenue  float64                `json:"revenue"`
	UserID   uint64                 `json:"user_id"`
	UserName string                 `json:"user_name"`
	From     MiteTime               `json:"from"`
	To       MiteTime               `json:"to"`
	Params   map[string]interface{} `json:"time_entries_params"`
}

type TimeEntry struct {
	ID       uint64   `json:"id"`
	Minutes  uint64   `json:"minutes"`
	DateAt   MiteTime `json:"date_at"`
	Note     string   `json:"note"`
	Billable bool     `json:"billable"`
	Locked   bool     `json:"locked"`
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

type GetTimeEntryResponseWrapper struct {
	TimeEntry *TimeEntry `json:"time_entry"`
}

type GetTimeEntriesGroupResponseWrapper struct {
	TimeEntry *TimeEntryGroup `json:"time_entry_group"`
}

// -------------------------------------------------------------
// ~ Get
// -------------------------------------------------------------

func (m *Mite) GetTimeEntriesGroup(from, to time.Time, filters map[string]string) ([]*TimeEntryGroup, error) {
	var timeGroupEntries []*GetTimeEntriesGroupResponseWrapper
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

func (m *Mite) GetTimeEntries(from, to time.Time, filters map[string]string) ([]*TimeEntry, error) {
	var timeResp []*GetTimeEntryResponseWrapper
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

func (m *Mite) GetTimeEntriesForProjectByService(from, to time.Time, projectID uint64) ([]*TimeEntryGroup, error) {
	params := map[string]string{
		ParamProjectID: fmt.Sprint(projectID),
		ParamGroupBy:   "service",
	}

	return m.GetTimeEntriesGroup(from, to, params)
}

func (m *Mite) GetTimeEntry(id string) (*TimeEntry, error) {
	var timeResp *GetTimeEntryResponseWrapper
	err := m.getAndDecodeFromSuffix("time_entry/"+id+".json", &timeResp, nil)
	if err != nil {
		return nil, err
	}
	return timeResp.TimeEntry, nil
}

// -------------------------------------------------------------
// ~ Create
// -------------------------------------------------------------

func (m *Mite) CreateTimeEntry(entry *TimeEntry) (*TimeEntry, error) {
	reqData := &GetTimeEntryResponseWrapper{TimeEntry: entry}
	resp, errRequest := m.postToMite("/time_entries.json", reqData)
	if errRequest != nil {
		return nil, errRequest
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Failed to create a time entry: " + resp.Status)
	}

	var respEntry = &GetTimeEntryResponseWrapper{}

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
func (m *Mite) UpdateTimeEntry(id string, update *TimeEntry) error {
	// Wrap time entry
	wrap := &GetTimeEntryResponseWrapper{TimeEntry: update}

	resp, errRequest := m.patchAtMite("/time_entry/"+id+".json", wrap)
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

func (m *Mite) DeleteTimeEntry(id string) error {
	resp, errRequest := m.deleteFromMite("/time_entry/"+id+".json", nil)
	if errRequest != nil {
		return errRequest
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("Failed to create a time entry: " + resp.Status)
	}
	return nil
}
