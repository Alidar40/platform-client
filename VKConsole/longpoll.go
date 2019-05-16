package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)



type LpResponse struct {
	Ts	int			`json:"ts"`
	Updates	[][]interface{}		`json:"updates"`
	Failed	int			`json:"failed"`
}



func QueryLongPollServer(server string, key string, ts int) (*LpResponse, error) {
	req, err := http.NewRequest("GET", "https://"+server, nil)

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
		return nil, err
	}
	defer resp.Body.Close()

	var lpResp LpResponse
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &lpResp)
	if err != nil {
		return nil, err
	}

	return &lpResp, nil
}
