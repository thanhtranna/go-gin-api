package go_gin_api

import (
	"encoding/json"

	"github.com/xinliangnote/go-gin-api/pkg/httpclient"

	"github.com/pkg/errors"
)

// interface address
var demoGetApi = "http://127.0.0.1:9999/demo/get/"

// Interface return structure
type demoGetResponse struct {
	Name string `json:"name"`
	Job  string `json:"job"`
}

// Initiate a request
func DemoGet(name string, opts ...httpclient.Option) (res *demoGetResponse, err error) {
	api := demoGetApi + name
	body, err := httpclient.Get(api, nil, opts...)
	if err != nil {
		return nil, err
	}

	res = new(demoGetResponse)
	err = json.Unmarshal(body, res)
	if err != nil {
		return nil, errors.Wrap(err, "DemoGet json unmarshal error")
	}

	return res, nil
}

// Set retry rules
func DemoGetRetryVerify(body []byte) (shouldRetry bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// Set alarm rules
func DemoGetAlarmVerify(body []byte) (shouldAlarm bool) {
	if len(body) == 0 {
		return true
	}

	return false
}

// Set up mock data
func DemoGetMock() (body []byte) {
	res := new(demoGetResponse)
	res.Name = "AA"
	res.Job = "AA_JOB"

	body, _ = json.Marshal(res)
	return body
}
