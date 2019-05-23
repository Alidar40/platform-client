package main

import(
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"

	"github.com/pkg/errors"
)

type GetUserResponse struct {
	Id		int	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
}

type VKError struct {
	ErrorCode	int	`json:"error_code"`
	ErrorMsg	string	`json:"error_msg"`
}

type GetUsersResponse struct {
	Error	 VKError		`json:"error,omitempty"`
	Response []GetUserResponse	`json:"response"`
}

func GetUserById(id int, token string, version float64) (*GetUserResponse, error) {
	req, err := http.NewRequest("GET", "https://api.vk.com/method/users.get", nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create a request to https://api.vk.com/method/users.get")
	}

	q := make(url.Values)
	q.Add("user_ids", fmt.Sprintf("%d", id))
	q.Add("access_token", token)
	q.Add("v", fmt.Sprintf("%.2f", version))
	req.URL.RawQuery = q.Encode()

	resp, err := vkc.Client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to query %q", req.URL)
	}
	defer resp.Body.Close()

	var getUsersResp GetUsersResponse
	err = json.NewDecoder(resp.Body).Decode(&getUsersResp)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal response from %q", req.URL)
	}

	if getUsersResp.Error.ErrorCode != 0 {
		return nil, errors.Errorf("failed to get user by id=%d with message '%s'", id, getUsersResp.Error.ErrorMsg)
	}

	return &getUsersResp.Response[0], nil
}
