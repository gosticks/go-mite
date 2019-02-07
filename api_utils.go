package mite

import (
	"fmt"
)

// GetCustomerByName returns a mite user
func (m *Mite) GetCustomerByName(name string) (*Customer, error) {
	var customersResponse []*getCustomersResponseWrapper

	params := map[string]string{
		"name": fmt.Sprint(name),
	}

	err := m.getAndDecodeFromSuffix("customers.json", &customersResponse, params)
	if err != nil {
		return nil, err
	}

	if len(customersResponse) > 1 {
		return nil, fmt.Errorf("customer name is not unique found %d customers mathing that name", len(customersResponse))
	}
	if len(customersResponse) == 0 {
		return nil, fmt.Errorf("Customer with name %s not found", name)
	}

	return customersResponse[0].Customer, nil
}

// GetProjectsForCustomer returns all projects for a customer
func (m *Mite) GetProjectsForCustomer(customerID uint64) ([]*Project, error) {
	var projectResponse []*getProjectsResponseWrapper

	params := map[string]string{
		"customer_id": fmt.Sprint(customerID),
	}

	err := m.getAndDecodeFromSuffix("projects.json", &projectResponse, params)
	if err != nil {
		return nil, err
	}

	projects := make([]*Project, len(projectResponse))

	// Unwrap all the data
	for i, p := range projectResponse {
		projects[i] = p.Project
		// fmt.Println("Project: ", p.Project.Name)
	}

	return projects, nil
}
