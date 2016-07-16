package connector

import (
	"bytes"
	"net/http"
)

func post(url string, b *bytes.Buffer, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, b)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	/*for k, v := range headers {
		req.Header.Add(k, v)
	}*/

	client := &http.Client{}
	return client.Do(req)
}

func get(url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	return client.Do(req)
}