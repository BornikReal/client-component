package service

import (
	"net/http"
)

func (k *KVService) Set(key string, value string) error {
	req, err := http.NewRequest(
		"POST", k.address, nil,
	)
	if err != nil {
		return err
	}
	req.URL.Query().Set("key", key)
	req.URL.Query().Set("value", value)
	resp, err := k.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
