package handlers

type handlers struct {
	client Client
}

func NewHandlers(client Client) *handlers {
	return &handlers{client: client}
}
