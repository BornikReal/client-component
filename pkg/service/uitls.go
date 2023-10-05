package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (k *KVService) getAddress(path string) string {
	return fmt.Sprintf("http://%s%s", k.address, path)
}

func answerFromString[T any](raw []byte) (T, error) {
	var ans T
	err := json.Unmarshal(raw, &ans)
	return ans, err
}

func (k *KVService) processOnlyQueryRequest(method string, path string, args map[string]string) ([]byte, error) {
	req, err := http.NewRequest(
		method, k.getAddress(path), nil,
	)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for key, value := range args {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := k.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		var ans httpAnswer
		ans, err = answerFromString[httpAnswer](resBody)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(ans.Message)
	}

	return resBody, nil
}
