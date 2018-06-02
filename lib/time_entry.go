package toggl

import (
	"encoding/json"
	"strconv"
)

type TimeEntry struct {
	At          string   `json:"at"`
	Billable    bool     `json:"billable"`
	Description string   `json:"description"`
	Duration    int64    `json:"duration"`
	Duronly     bool     `json:"duronly"`
	ID          int      `json:"id"`
	Start       string   `json:"start"`
	Tags        []string `json:"tags"`
	UID         int      `json:"uid"`
	PID         int      `json:"pid"`
	WID         int      `json:"wid"`
}

type CurrentResponse struct {
	Data TimeEntry `json:"data"`
}

func (timeEntry TimeEntry) AddParam() interface{} {
	param := make(map[string]map[string]interface{})
	param["time_entry"] = make(map[string]interface{})
	if timeEntry.PID != 0 {
		param["time_entry"]["pid"] = timeEntry.PID
	}
	param["time_entry"]["wid"] = timeEntry.WID
	param["time_entry"]["description"] = timeEntry.Description
	param["time_entry"]["created_with"] = "sachaos/toggl"
	return param
}

func (cl *Client) GetCurrentTimeEntry() (CurrentResponse, error) {
	var response CurrentResponse

	res, err := cl.do("GET", "time_entries/current", nil)
	if err != nil {
		return CurrentResponse{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&response); err != nil {
		return CurrentResponse{}, err
	}

	return response, nil
}

func (cl *Client) PostStartTimeEntry(timeEntry TimeEntry) (response CurrentResponse, err error) {
	res, err := cl.do("POST", "time_entries/start", timeEntry.AddParam())
	if err != nil {
		return CurrentResponse{}, err
	}

	enc := json.NewDecoder(res.Body)
	if err := enc.Decode(&response); err != nil {
		return CurrentResponse{}, err
	}

	return response, nil
}

func (cl *Client) PutStopTimeEntry(id int) error {
	id_string := strconv.Itoa(id)

	_, err := cl.do("PUT", "time_entries/"+id_string+"/stop", nil)

	return err
}
