package toggl

import (
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/franela/goreq"
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
	Wid         int      `json:"wid"`
}

type CurrentResponse struct {
	Data TimeEntry `json:"data"`
}

func (timeEntry TimeEntry) AddParam() interface{} {
	param := make(map[string]map[string]interface{})
	param["time_entry"] = make(map[string]interface{})
	param["time_entry"]["description"] = timeEntry.Description
	param["time_entry"]["created_with"] = "sachaos/toggl"
	return param
}

func GetCurrentTimeEntry(token string) (CurrentResponse, error) {
	var response CurrentResponse
	basic := base64.StdEncoding.EncodeToString([]byte(token + ":api_token"))
	res, err := goreq.Request{
		Method:      "GET",
		Uri:         "https://www.toggl.com/api/v8/time_entries/current",
		ContentType: "application/json",
	}.WithHeader("Authorization", "Basic "+basic).Do()
	if err != nil {
		return CurrentResponse{}, err
	}
	res.Body.FromJsonTo(&response)
	return response, nil
}

func PostStartTimeEntry(timeEntry TimeEntry, token string) error {
	basic := base64.StdEncoding.EncodeToString([]byte(token + ":api_token"))
	body_text, err := json.Marshal(timeEntry.AddParam())
	if err != nil {
		return err
	}
	_, err = goreq.Request{
		Method:      "POST",
		Uri:         "https://www.toggl.com/api/v8/time_entries/start",
		ContentType: "application/json",
		Body:        string(body_text),
	}.WithHeader("Authorization", "Basic "+basic).Do()
	if err != nil {
		return err
	}
	return nil
}

func PutStopTimeEntry(id int, token string) error {
	basic := base64.StdEncoding.EncodeToString([]byte(token + ":api_token"))
	id_string := strconv.Itoa(id)
	_, err := goreq.Request{
		Method:      "PUT",
		Uri:         "https://www.toggl.com/api/v8/time_entries/" + id_string + "/stop",
		ContentType: "application/json",
	}.WithHeader("Authorization", "Basic "+basic).Do()
	if err != nil {
		return err
	}
	return nil
}
