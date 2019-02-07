package mite

import "time"

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

// Account is a type for mite the account
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type accountResponseWrapper struct {
	Account *Account `json:"account"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

// GetAccount returns the current mite account
func (m *Mite) GetAccount() (*Account, error) {
	var resp *accountResponseWrapper
	err := m.getAndDecodeFromSuffix("account.json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.Account, nil
}

// GetMyself returns the current user
func (m *Mite) GetMyself() (*User, error) {
	var resp *getUsersResponseWrapper
	err := m.getAndDecodeFromSuffix("myself.json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}
