package toggl

import (
	"errors"
)

type Workspace struct {
	Admin                       bool   `json:"admin"`
	APIToken                    string `json:"api_token"`
	At                          string `json:"at"`
	DefaultCurrency             string `json:"default_currency"`
	DefaultHourlyRate           int    `json:"default_hourly_rate"`
	IcalEnabled                 bool   `json:"ical_enabled"`
	ID                          int    `json:"id"`
	Name                        string `json:"name"`
	OnlyAdminsMayCreateProjects bool   `json:"only_admins_may_create_projects"`
	OnlyAdminsSeeBillableRates  bool   `json:"only_admins_see_billable_rates"`
	OnlyAdminsSeeTeamDashboard  bool   `json:"only_admins_see_team_dashboard"`
	Premium                     bool   `json:"premium"`
	Profile                     int    `json:"profile"`
	ProjectsBillableByDefault   bool   `json:"projects_billable_by_default"`
	Rounding                    int    `json:"rounding"`
	RoundingMinutes             int    `json:"rounding_minutes"`
	Subscription                struct {
		CreatedAt            string      `json:"created_at"`
		DeletedAt            interface{} `json:"deleted_at"`
		Description          string      `json:"description"`
		UpdatedAt            interface{} `json:"updated_at"`
		VatApplicable        bool        `json:"vat_applicable"`
		VatInvalidAcceptedAt interface{} `json:"vat_invalid_accepted_at"`
		VatInvalidAcceptedBy interface{} `json:"vat_invalid_accepted_by"`
		VatValid             bool        `json:"vat_valid"`
		VatValidatedAt       interface{} `json:"vat_validated_at"`
		WorkspaceID          int         `json:"workspace_id"`
	} `json:"subscription"`
}

type Workspaces []Workspace

func (repository Workspaces) FindByID(id int) (*Workspace, error) {
	for _, item := range repository {
		if item.ID == id {
			return &item, nil
		}
	}
	return nil, errors.New("Find Failed")
}

func FetchWorkspaces(token string) (Workspaces, error) {
	var workspaces Workspaces
	res, err := Request("GET", "/workspaces", nil, token)
	if err != nil {
		return Workspaces{}, err
	}
	err = res.Body.FromJsonTo(&workspaces)
	if err != nil {
		return Workspaces{}, err
	}
	return workspaces, nil
}
