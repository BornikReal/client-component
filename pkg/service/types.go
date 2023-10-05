package service

const getPath = "/kv"
const setPath = "/kv"

type httpAnswer struct {
	Code    int64  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type getAnswer struct {
	Value string `json:"value,omitempty"`
}
