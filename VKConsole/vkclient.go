package main

import(
	"net/http"
	"net/http/cookiejar"
	"time"
	"os"
	"math"
	"fmt"

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

func (vkc *VKClient) ListenToLongPollServer() {
	for {
		lpResp, err := QueryLongPollServer(vkc.Server, vkc.Key, vkc.Ts)
		if err != nil {
			fmt.Printf("FATAL: %+v\n", err)
			os.Exit(1)
		}

		vkc.Ts = lpResp.Ts

		switch (lpResp.Failed) {
			case 1:
				continue
			case 2:
				fmt.Println("key timeout")
				os.Exit(2)
				//See point 1 in main.go
			case 3:
				fmt.Println("key and ts timeout")
				os.Exit(2)
				//See point 1 in main.go
			case 4:
				fmt.Println("version is incorrect")
				os.Exit(2)
		}

		for _, update := range lpResp.Updates {
			switch (update[0].(float64)) {
				case 4:
					userId := int(math.Abs(update[3].(float64)))
					title := update[5].(string)
					if (userId - 2000000000) > 0 {
						fmt.Println("New message in one of your chats")
						fmt.Println("\tIt says: \"" + title + "\"")
						break
					}

					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Printf("FATAL: %+v\n", err)
						os.Exit(1)
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					fmt.Println("New message in conversation with " + userName)
					fmt.Println("\tIt says: \"" + title + "\"")
					break
				case 8:
					userId := int(math.Abs(update[1].(float64)))
					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Printf("FATAL: %+v\n", err)
						os.Exit(1)
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					fmt.Println("Your friend " + userName + " became online")
					break
				case 9:
					userId := int(math.Abs(update[1].(float64)))
					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Printf("FATAL: %+v\n", err)
						os.Exit(1)
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					fmt.Println("Your friend " + userName + " became offline")
					break
			}
		}
	}
}
