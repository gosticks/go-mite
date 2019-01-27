package mite

import (
	"path"
)

// -------------------------------------------------------------
// ~ Const
// -------------------------------------------------------------

const (
	MiteURL        = "mite.yo.lk"
	MiteTimeFormat = "2006-01-02"
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

type ServiceHourlyRates struct {
	ServiceID  uint64 `json:"service_id"`
	HourlyRate uint64 `json:"hourly_rate"`
}

type Mite struct {
	// l        *zap.SugaredLogger
	Prefix   string
	Username string
	ApiKey   string
	AppName  string
}

// -------------------------------------------------------------
// ~ Functions
// -------------------------------------------------------------

func NewMiteAPI(username, team, apiKey, appName string) *Mite {
	return &Mite{
		Prefix:   team,
		Username: username,
		ApiKey:   apiKey,
		AppName:  appName,
	}
}

func (m *Mite) GetMitePath() string {
	return "https://" + m.Prefix + "." + MiteURL
}

func (m *Mite) MitePathWithParam(suffix string) string {
	return "https://" + path.Join((m.Prefix+"."+MiteURL), suffix)
}
