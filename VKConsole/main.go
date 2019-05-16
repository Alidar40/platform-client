package main

import(
	"fmt"
	"math"
)

var vkc *VKClient

func main() {
	//1. Auth
	/* Well, this part is not so easy as you may think
	Check this out:
		https://vk.com/dev.php?method=messages_api
	*/

	//2. Establishing connection to the long poll server
	ci, err := readConnectionInfo("connectioninfo.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	vkc = NewVKClient(ci.Timeout, ci.Token, ci.Server, ci.Key, ci.Ts)

	//3. Listening
	for {
		lpResp, err := QueryLongPollServer(vkc.Server, vkc.Key, vkc.Ts)
		if err != nil {
			fmt.Println(err)
			return
		}

		vkc.Ts = lpResp.Ts

		switch (lpResp.Failed) {
			case 1:
				continue
			case 2:
				fmt.Println("key timeout")
				return
				//See point 1
			case 3:
				fmt.Println("key and ts timeout")
				return
				//See point 1
			case 4:
				fmt.Println("version is incorrect")
				return
		}

		for _, update := range lpResp.Updates {
			switch (update[0].(float64)) {
				case 4:
					userId := int(math.Abs(update[3].(float64)))
					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Println(err)
						return
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					title := update[5].(string)
					fmt.Println("New message from " + userName)
					fmt.Println("It says: \"" + title + "\"")
					break
				case 8:
					userId := int(math.Abs(update[1].(float64)))
					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Println(err)
						return
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					fmt.Println("Your friend " + userName + " became online")
					break
				case 9:
					userId := int(math.Abs(update[1].(float64)))
					getUserResp, err := GetUserById(userId, vkc.Token, 5.95)
					if err != nil {
						fmt.Println(err)
						return
					}

					userName := getUserResp.FirstName + " "  + getUserResp.LastName
					fmt.Println("Your friend " + userName + " became offline")
					break
				default:
			}
		}
	}
}
