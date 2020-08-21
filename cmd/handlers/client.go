package handlers

import "net/http"

//go:generate mockgen -package=mock -destination=./mock/client.go -source=client.go
type Client interface {
	Get(string) (*http.Response, error)
}
