package logit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"logClient/client/http"
	ta "logClient/types/user"
	"logClient/utils"
	// Http "net/http"
)

const (
	getUser    = baseUrl + user
	updateUser = baseUrl + user
)

var GetUser = func(tk string) ([]byte, error) {

	client := http.CreateClient(getUser, map[string]string{
		"Authorization": "Bearer " + tk,
		"Content-Type":  "application/json",
	})
	req, err := client.MakeRequest(utils.GET, nil)

	if err != nil {
		// fmt.Println(err)
		return []byte{}, err
	}
	response, err := client.Act(req)
	if err != nil {
		fmt.Println(err)
		return []byte{}, err
	}

	return response, nil
}

var UpdateUser = func(at *ta.User, tk string) ([]byte, error) {
	client := http.CreateClient(updateUser, map[string]string{
		"Authorization": "Bearer " + tk,
		"Content-Type":  "application/json",
	})
	raw, err := json.Marshal(at)
	if err != nil {
		return []byte{}, err
	}
	req, err := client.MakeRequest(utils.PUT, bytes.NewBuffer(raw))

	if err != nil {
		return []byte{}, err
	}
	response, err := client.Act(req)
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}
