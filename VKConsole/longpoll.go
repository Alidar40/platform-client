package main

import(
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/pkg/errors"
)



type LpResponse struct {
	Ts	int			`json:"ts"`
	Updates	[][]interface{}		`json:"updates"`
	Failed	int			`json:"failed"`
}



func QueryLongPollServer(server string, key string, ts int) (*LpResponse, error) {
	req, err := http.NewRequest("GET", "https://"+server, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to created request to https://%q", server)
	}

	q := req.URL.Query()
	q.Add("act", "a_check")
	q.Add("key", key)
	q.Add("ts", fmt.Sprintf("%d", vkc.Ts))
	q.Add("wait", "25")
	q.Add("mode", "2")
	q.Add("version", "3")
	req.URL.RawQuery = q.Encode()

	resp, err := vkc.Client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query %q", req.URL)
	}
	defer resp.Body.Close()

	var lpResp LpResponse
	err = json.NewDecoder(resp.Body).Decode(&lpResp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response from %q", req.URL)
	}

	return &lpResp, nil
}
