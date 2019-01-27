package mite

import (
	"fmt"
	"time"
)

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

type Customer struct {
	ID                    uint64               `json:"id"`
	Name                  string               `json:"name"`
	Note                  string               `json:"note"`
	ActiveHourlyRate      string               `json:"active_hourly_rate"`
	HourlyRate            uint64               `json:"hourly_rate"`
	Archived              bool                 `json:"archived"`
	HourlyRatesPerService []ServiceHourlyRates `json:"hourly_rates_per_service"`
	CreatedAt             time.Time            `json:"created_at"`
	UpdatedAt             time.Time            `json:"updated_at"`
}

func (c *Customer) String() string {
	return fmt.Sprintf("%d: %s for %s (archived: %t)", c.ID, c.Name, c.Name, c.Archived)
}

type GetCustomersResponseWrapper struct {
	Customer *Customer `json:"customer"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

func (m *Mite) GetAllCustomers() ([]*Customer, error) {
	var cusRes []*GetCustomersResponseWrapper
	err := m.getAndDecodeFromSuffix("customers.json", &cusRes, nil)
	if err != nil {
		return nil, err
	}

	customers := make([]*Customer, len(cusRes))

	// Unwrap all the data
	for i, c := range cusRes {
		customers[i] = c.Customer
		//fmt.Println("Customer", c.Customer)
	}

	return customers, nil
}
