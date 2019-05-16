package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type GetUserResponse struct {
	Id		int	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
}

type GetUsersResponse struct {
	Response []GetUserResponse	`json:"response"`
}

func GetUserById(id int, token string, version float64) (*GetUserResponse, error) {
	req, err := http.NewRequest("GET", "https://api.vk.com/method/users.get", nil)

	q := req.URL.Query()
	q.Add("user_ids", fmt.Sprintf("%d", id))//strconv.Itoa(id))
	q.Add("access_token", token)
	q.Add("v", fmt.Sprintf("%.2f", version))
	req.URL.RawQuery = q.Encode()

	resp, err := vkc.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var getUsersResp GetUsersResponse
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBody, &getUsersResp)
	if err != nil {
		return nil, err
	}

	return &getUsersResp.Response[0], nil
}
