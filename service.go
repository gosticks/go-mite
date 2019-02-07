package mite

import (
	"fmt"
	"strconv"
	"time"
)

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

// Service mite object
type Service struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Note       string    `json:"note"`
	HourlyRate uint64    `json:"hourly_rate"`
	Archived   bool      `json:"archived"`
	Billable   bool      `json:"billable"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (s *Service) String() string {
	return fmt.Sprintf("%d: %s (archived: %t)", s.ID, s.Name, s.Archived)
}

type getServicesResponseWrapper struct {
	Service *Service `json:"service"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

// GetServices returns filters services from mite
func (m *Mite) GetServices(filters map[string]string) ([]*Service, error) {
	var serviceResp []*getServicesResponseWrapper
	err := m.getAndDecodeFromSuffix("services.json", &serviceResp, filters)
	if err != nil {
		return nil, err
	}

	services := make([]*Service, len(serviceResp))

	// Unwrap all the data
	for i, s := range serviceResp {
		services[i] = s.Service
		// fmt.Println("Service: ", s.Service)
	}

	return services, nil
}

// GetService returns a service for a id
func (m *Mite) GetService(id uint64) (*Service, error) {
	var resp *getServicesResponseWrapper
	err := m.getAndDecodeFromSuffix("service"+strconv.FormatUint(id, 10)+".json", &resp, nil)
	if err != nil {
		return nil, err
	}
	return resp.Service, nil
}
