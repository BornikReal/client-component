package service

import "net/http"

type KVService struct {
	client  http.Client
	address string
}

func NewKVService(client http.Client, address string) *KVService {
	return &KVService{
		client:  client,
		address: address,
	}
}
