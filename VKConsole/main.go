package main

import(
	"fmt"
	"os"
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
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	vkc, err = NewVKClient(ci.Timeout, ci.Token, ci.Server, ci.Key, ci.Ts)
	if err != nil {
		fmt.Printf("FATAL: %+v\n", err)
		os.Exit(1)
	}

	//3. Listening
	vkc.ListenToLongPollServer()
}
