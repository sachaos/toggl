package toggl

import (
	"encoding/base64"

	"github.com/franela/goreq"
)

type TimeEntry struct {
	At          string   `json:"at"`
	Billable    bool     `json:"billable"`
	Description string   `json:"description"`
	Duration    int      `json:"duration"`
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

func FetchCurrent(token string) (CurrentResponse, error) {
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
