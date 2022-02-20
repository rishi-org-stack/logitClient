package logit

import (
	"bytes"
	"encoding/json"
	"logClient/client/http"
	ta "logClient/types/auth"
	"logClient/utils"
)

const (
	getToken = baseUrl + auth
)

var GetToken = func(at *ta.AuthRequest) ([]byte, error) {
	client := http.CreateClient(getToken, map[string]string{})
	raw, err := json.Marshal(at)
	if err != nil {
		return []byte{}, err
	}
	req, err := client.MakeRequest(utils.POST, bytes.NewBuffer(raw))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		return []byte{}, err
	}
	response, err := client.Act(req)
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}
