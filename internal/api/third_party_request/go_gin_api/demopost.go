package go_gin_api

import (
	"encoding/json"
	"net/url"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

// interface address
var demoPostApi = "http://127.0.0.1:9999/demo/post/"

// Interface return structure
type demoPostResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// Initiate a request
func DemoPost(name string, opts ...httpclient.Option) (res *demoPostResponse, err error) {
	api := demoPostApi
	params := url.Values{}
	params.Set("name", name)
	body, err := httpclient.PostForm(api, params, opts...)
	if err != nil {
		return nil, err
	}

	res = new(demoPostResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, errors.Wrap(err, "DemoPost json unmarshal error")
	}

	return res, nil
}

// Set retry rules
func DemoPostRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// Set alarm rules
func DemoPostAlarmVerify(body []byte) (shouldAlarm bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// Set up mock data
func DemoPostMock() (body []byte) {
	res := new(demoPostResponse)
	res.Name = "BB"
	res.Job = "BB_JOB"

	body, _ = json.Marshal(res)
	return body
}
