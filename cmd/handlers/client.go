package handlers

import "net/http"

type Client interface {
	Get(string) (*http.Response, error)
}
