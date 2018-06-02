package toggl

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const (
	baseURI    = "https://www.toggl.com/api/v8"
	retryCount = 3
)

var (
	retryCountExceeded = errors.New("API request retry count exceeded")
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(client *http.Client, token string) *Client {
	return &Client{
		client: client,
		token:  token,
	}
}

func NewDefaultClient(token string) *Client {
	return &Client{
		client: http.DefaultClient,
		token:  token,
	}
}

func (cl *Client) do(method string, endpoint string, param interface{}) (res *http.Response, err error) {
	uri, _ := url.Parse(baseURI)
	uri.Path = path.Join(uri.Path, endpoint)

	req, err := http.NewRequest(method, uri.String(), nil)
	if err != nil {
		return
	}

	basic := base64.StdEncoding.EncodeToString([]byte(cl.token + ":api_token"))
	req.Header.Add("Authorization", "Basic "+basic)

	var buff []byte
	if param != nil {
		buff, err = json.Marshal(param)
		if err != nil {
			return
		}
	}
	req.Body = ioutil.NopCloser(bytes.NewReader(buff))

	count := 0
	for count < retryCount {
		res, err := cl.client.Do(req)
		if err == nil {
			return res, err
		}
		count++
	}

	return nil, retryCountExceeded
}
