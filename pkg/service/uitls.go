package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const getPath = "/kv"
const setPath = "/kv"

func (k *KVService) getAddress(path string) string {
	return fmt.Sprintf("http://%s%s", k.address, path)
}

type httpAnswer struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}

func httpAnswerFromString(raw []byte) (httpAnswer, error) {
	var ans httpAnswer
	err := json.Unmarshal(raw, &ans)
	return ans, err
}

func (k *KVService) processOnlyQueryRequest(method string, path string, args map[string]string) (string, error) {
	req, err := http.NewRequest(
		method, k.getAddress(path), nil,
	)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	for key, value := range args {
		q.Set(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := k.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ans, err := httpAnswerFromString(resBody)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New(ans.Message)
	}

	return ans.Value, nil
}
