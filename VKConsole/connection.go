package main

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
)

type ConnectionInfo struct {
	Timeout	int
	Token	string
	Server	string
	Key	string
	Ts	int
}

func readConnectionInfo(filename string) (*ConnectionInfo, error) {
	connectionInfoFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer connectionInfoFile.Close()

	connectionInfoBytes, _ := ioutil.ReadAll(connectionInfoFile)

	var connectionInfo ConnectionInfo
	json.Unmarshal(connectionInfoBytes, &connectionInfo)

	return &connectionInfo, nil
}
