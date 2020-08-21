package handlers

import (
	"net/http"
	"sync"
)

type handlers struct {
	client      Client
	initialised sync.Once
}

func NewHandlers() *handlers {
	return &handlers{}
}

func (h *handlers) SetClient(client Client) {
	h.client = client
}

func (h *handlers) getClient() Client {
	h.initialised.Do(h.init)
	return h.client
}

func (h *handlers) init() {
	if h.client == nil {
		h.client = http.DefaultClient
	}
}
