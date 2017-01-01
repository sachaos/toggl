package toggl

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"path"

	"github.com/franela/goreq"
)

const (
	baseURI    = "https://www.toggl.com/api/v8"
	retryCount = 3
)

func Request(method string, endpoint string, param interface{}, token string) (*goreq.Response, error) {
	// format endpoint
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	// format param
	var bodyText []byte
	var err error
	if param != nil {
		bodyText, err = json.Marshal(param)
		if err != nil {
			return nil, err
		}
	}

	// format token
	basic := base64.StdEncoding.EncodeToString([]byte(token + ":api_token"))

	count := 0
	var response *goreq.Response
	for count < retryCount {
		response, err = goreq.Request{
			Method:      method,
			Uri:         uri.String(),
			ContentType: "application/json",
			Body:        string(bodyText),
		}.WithHeader("Authorization", "Basic "+basic).Do()
		if err == nil {
			return response, err
		}
		count++
	}
	return response, err
}
