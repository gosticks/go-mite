package mite

import (
	"path"
)

// -------------------------------------------------------------
// ~ Const
// -------------------------------------------------------------

const (
	// MiteURL is the current mite url
	MiteURL = "mite.yo.lk"
	// TimeFormat is the time format used by mite
	TimeFormat = "2006-01-02"
)

// filter options
const (
	ParamProjectID  = "project_id"
	ParamCustomerID = "customer_id"
	ParamGroupBy    = "group_by"
	ParamFrom       = "from"
	ParamTo         = "to"
)

// -------------------------------------------------------------
// ~ Types
// -------------------------------------------------------------

// ServiceHourlyRates mite object
type ServiceHourlyRates struct {
	ServiceID  uint64 `json:"service_id"`
	HourlyRate uint64 `json:"hourly_rate"`
}

// Mite is the interface used for the api
type Mite struct {
	// l        *zap.SugaredLogger
	Prefix   string
	Username string
	APIKey   string
	AppName  string
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

// NewMiteAPI creates a new mite api struct
func NewMiteAPI(username, team, apiKey, appName string) *Mite {
	return &Mite{
		Prefix:   team,
		Username: username,
		APIKey:   apiKey,
		AppName:  appName,
	}
}

// GetMitePath returns a mite path for the for the current workspace
func (m *Mite) GetMitePath() string {
	return "https://" + m.Prefix + "." + MiteURL
}

func (m *Mite) mitePathWithParam(suffix string) string {
	return "https://" + path.Join((m.Prefix+"."+MiteURL), suffix)
}
