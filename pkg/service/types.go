package service

const getPath = "/kv"
const setPath = "/kv"

type httpAnswer struct {
	Message string `json:"message,omitempty"`
}

type getAnswer struct {
	Value string `json:"value,omitempty"`
}
