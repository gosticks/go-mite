package mite

import (
	"fmt"
	"strconv"
	"time"
)

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

// User mite object
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

type getUsersResponseWrapper struct {
	User *User `json:"user"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

// GetUsers returns all users in a mite workspace
func (m *Mite) GetUsers() ([]*User, error) {
	var usersResponse []*getUsersResponseWrapper
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

// GetUser returns a mite user for id
func (m *Mite) GetUser(id uint64) (*User, error) {
	var resp *getUsersResponseWrapper
	err := m.getAndDecodeFromSuffix("users/"+strconv.FormatUint(id, 10)+".json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}
