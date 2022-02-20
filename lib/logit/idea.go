package logit

import (
	"bytes"
	"encoding/json"
	"logClient/client/http"
	ta "logClient/types/idea"
	"logClient/utils"
)

const (
	ideaUrl = baseUrl + idea
)

var AddIdea = func(idea *ta.Idea, tk string) ([]byte, error) {
	client := http.CreateClient(ideaUrl, map[string]string{
		"Authorization": "Bearer " + tk,
		"Content-Type":  "application/json",
	})
	raw, err := json.Marshal(idea)
	if err != nil {
		return []byte{}, err
	}
	req, err := client.MakeRequest(utils.POST, bytes.NewBuffer(raw))

	if err != nil {
		return []byte{}, err
	}
	response, err := client.Act(req)
	if err != nil {
		return []byte{}, err
	}
	return response, nil
}
