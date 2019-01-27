package mite

import (
	"errors"
	"fmt"
)

func (m *Mite) GetCustomerByName(name string) (*Customer, error) {
	customers, err := m.GetAllCustomers()
	if err != nil {
		return nil, err
	}

	for _, c := range customers {
		if c.Name == name {
			return c, nil
		}
	}

	return nil, errors.New("Could not find customer")
}

func (m *Mite) GetProjectsForCustomer(customerID uint64) ([]*Project, error) {
	var projectResponse []*GetProjectsResponseWrapper

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
