package service

import (
	"io"
	"net/http"
)

func (k *KVService) Get(key string) (string, error) {
	req, err := http.NewRequest(
		"GET", k.address, nil,
	)
	if err != nil {
		return "", err
	}
	req.URL.Query().Set("key", key)
	resp, err := k.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
