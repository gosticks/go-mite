package mite

import (
	"fmt"
	"time"
)

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Archived  bool      `json:"archived"`
	Language  string    `json:"language"`
	Role      string    `json:"role"`
}

func (u *User) String() string {
	return fmt.Sprintf("%s, %s (archived: %t, id: %d, role: %s)", u.Name, u.Email, u.Archived, u.ID, u.Role)
}

type GetUsersResponseWrapper struct {
	User *User `json:"user"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

func (m *Mite) GetUsers() ([]*User, error) {
	var usersResponse []*GetUsersResponseWrapper
	err := m.getAndDecodeFromSuffix("users.json", &usersResponse, nil)
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(usersResponse))

	// Unwrap all the data
	for i, u := range usersResponse {
		users[i] = u.User
		// fmt.Println(u.User.Name, u.User.Email)
	}

	return users, nil
}

func (m *Mite) GetUser(id string) (*User, error) {
	var resp *GetUsersResponseWrapper
	err := m.getAndDecodeFromSuffix("users/"+id+".json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}