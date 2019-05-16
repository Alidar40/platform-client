package main

import(
	"net/http"
	"net/http/cookiejar"
	"time"
)

type VKClient struct {
	Client	*http.Client
	Token	string
	Server	string
	Key	string
	Ts	int
}

func NewVKClient(timeout int, token string, server string, key string, ts int) (*VKClient) {
	vkc := new(VKClient)

	cookieJar, _ := cookiejar.New(nil)
	vkc.Client = &http.Client{
		Jar: cookieJar,
		Timeout: time.Duration(timeout) * time.Second,
	}
	vkc.Token = token
	vkc.Server = server
	vkc.Key = key
	vkc.Ts = ts

	return vkc
}
