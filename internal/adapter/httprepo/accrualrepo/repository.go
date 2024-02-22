package accrualrepo

import (
	"net/http"
)

type Repository struct {
	client  *http.Client
	baseURL string
}

func NewRepository(client *http.Client, baseURL string) *Repository {
	return &Repository{
		client:  client,
		baseURL: baseURL,
	}
}
