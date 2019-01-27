package mite

import "time"

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountResponseWrapper struct {
	Account *Account `json:"account"`
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

func (m *Mite) GetAccount() (*Account, error) {
	var resp *AccountResponseWrapper
	err := m.getAndDecodeFromSuffix("account.json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.Account, nil
}

func (m *Mite) GetMyself() (*User, error) {
	var resp *GetUsersResponseWrapper
	err := m.getAndDecodeFromSuffix("myself.json", &resp, nil)
	if err != nil {
		return nil, err
	}

	return resp.User, nil
}
