package toggl

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Project struct {
	Active        bool   `json:"active"`
	ActualHours   int    `json:"actual_hours"`
	At            string `json:"at"`
	AutoEstimates bool   `json:"auto_estimates"`
	Billable      bool   `json:"billable"`
	Color         string `json:"color"`
	CreatedAt     string `json:"created_at"`
	HexColor      string `json:"hex_color"`
	ID            int    `json:"id"`
	IsPrivate     bool   `json:"is_private"`
	Name          string `json:"name"`
	Template      bool   `json:"template"`
	Wid           int    `json:"wid"`
}

type Projects []Project

func (repository Projects) FindByID(id int) (Project, error) {
	for _, item := range repository {
		if item.ID == id {
			return item, nil
		}
	}
	return Project{}, errors.New("Find Failed")
}

func (cl *Client) FetchWorkspaceProjects(wid int) (Projects, error) {
	var projects Projects

	res, err := cl.do("GET", "/workspaces/"+strconv.Itoa(wid)+"/projects", nil)
	if err != nil {
		return Projects{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&projects); err != nil {
		return Projects{}, err
	}

	return projects, nil
}
