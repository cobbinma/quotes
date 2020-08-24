package handlers

import "net/http"

type handlers struct {
	client Client
}

type handlerOptions struct {
	client Client
}

func WithCustomClient(client Client) func(*handlerOptions) *handlerOptions {
	return func(ho *handlerOptions) *handlerOptions {
		ho.client = client
		return ho
	}
}

func NewHandlers(options ...func(*handlerOptions) *handlerOptions) *handlers {
	opts := &handlerOptions{
		client: http.DefaultClient,
	}
	for i := range options {
		opts = options[i](opts)
	}
	return &handlers{
		client: opts.client,
	}
}
