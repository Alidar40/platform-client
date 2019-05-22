package main

import(
	"io/ioutil"
	"encoding/json"
	"os"

	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(err, "failed to open %q", filename)
	}
	defer connectionInfoFile.Close()

	connectionInfoBytes, err := ioutil.ReadAll(connectionInfoFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %q", filename)
	}

	var connectionInfo ConnectionInfo
	err = json.Unmarshal(connectionInfoBytes, &connectionInfo)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %q", filename)
	}

	return &connectionInfo, nil
}
