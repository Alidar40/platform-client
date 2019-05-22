package main

import(
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/pkg/errors"
)

type VKClient struct {
	Client	*http.Client
	Token	string
	Server	string
	Key	string
	Ts	int
}

func NewVKClient(timeout int, token string, server string, key string, ts int) (*VKClient, error) {
	vkc := new(VKClient)

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cookiejar for vkclient")
	}

	vkc.Client = &http.Client{
		Jar: cookieJar,
		Timeout: time.Duration(timeout) * time.Second,
	}
	vkc.Token = token
	vkc.Server = server
	vkc.Key = key
	vkc.Ts = ts

	return vkc, nil
}
