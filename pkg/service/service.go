package service

import "net/http"

type ClientComponent struct {
	client  *http.Client
	address string
}

func NewKVService(client *http.Client, address string) *ClientComponent {
	return &ClientComponent{
		client:  client,
		address: address,
	}
}
