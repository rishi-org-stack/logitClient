package http

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	url     string
	headers map[string]string
}

func CreateClient(url string, headers map[string]string) *HttpClient {
	return &HttpClient{
		url:     url,
		headers: headers,
	}
}

func (c *HttpClient) MakeRequest(method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.url, body)
	for key, val := range c.headers {
		req.Header.Add(key, val)
	}
	// req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	return req, err
}

func (c *HttpClient) Act(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	fmt.Println("=========================Request Sent========================")
	fmt.Println("url: ", req.URL)
	fmt.Println("method", req.Method)
	return raw, nil
}
