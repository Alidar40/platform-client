package main

import(
	"fmt"
	"net/http"
	//"net/url"
	"net/http/cookiejar"
	//"golang.org/x/net/html"
	//"strings"
	//"io/ioutil"
	"encoding/json"
	"strconv"
	"time"
)

type Update struct {
	Code	int
	Text	string
	UserId	string
}

type LpResponse struct {
	Ts	int	`json:"ts"`
	//Updates	[]Update	`json:"updates,omitempty"`
}

func QueryLongPollServer(server string, key string, ts int) (int, error) {
	resp, err := client.Get("https://"+server+"?act=a_check&key="+key+"&ts="+strconv.Itoa(ts)+"wait=25&mode=2&version=2")
	if err != nil {
		return ts, err
	}
	defer resp.Body.Close()

	var lpResp LpResponse
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&lpResp)
	if err != nil {
		return ts, err
	}

	fmt.Println(lpResp.Ts)
	return lpResp.Ts, nil
}


var client *http.Client

func main() {
	cookieJar, _ := cookiejar.New(nil)
	tr := &http.Transport{
		IdleConnTimeout:    30 * time.Second,
		ExpectContinueTimeout: 30 * time.Second,
	}
	client = &http.Client{
		Transport: tr,
		Jar: cookieJar,
		Timeout: 30 * time.Second,
	}

	server := "SERVER"
	key := "KEY"
	ts := 111111111
	for {
		var err error
		ts, err = QueryLongPollServer(server, key, ts)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
