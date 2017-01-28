package toggl

import (
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

func GetCurrentTimeEntry(token string) (CurrentResponse, error) {
	var response CurrentResponse

	res, err := Request("GET", "time_entries/current", nil, token)

	if err != nil {
		return CurrentResponse{}, err
	}
	res.Body.FromJsonTo(&response)
	return response, nil
}

func PostStartTimeEntry(timeEntry TimeEntry, token string) (response CurrentResponse, err error) {
	res, err := Request("POST", "time_entries/start", timeEntry.AddParam(), token)
	if err != nil {
		return CurrentResponse{}, err
	}
	res.Body.FromJsonTo(&response)

	if err != nil {
		return
	}
	return response, nil
}

func PutStopTimeEntry(id int, token string) error {
	id_string := strconv.Itoa(id)

	_, err := Request("PUT", "time_entries/"+id_string+"/stop", nil, token)

	if err != nil {
		return err
	}
	return nil
}
