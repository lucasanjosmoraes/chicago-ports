package http

import (
	"bytes"
	"context"
)

// Header represents a HTTP header.
type Header struct {
	Key   string
	Value string
}

// ClientResponse represents methods result.
type ClientResponse struct {
	StatusCode int
	Body       []byte
}

// Client defines everything that an adapter needs to provide to every http client.
type Client interface {
	Delete(ctx context.Context, url string, headers []Header) (ClientResponse, error)
	Get(ctx context.Context, url string, headers []Header) (ClientResponse, error)
	Patch(ctx context.Context, url string, body []byte, headers []Header) (ClientResponse, error)
	Post(ctx context.Context, url string, body []byte, headers []Header) (ClientResponse, error)
	PostForm(ctx context.Context, url string, body *bytes.Buffer, formHeader string, headers []Header) (ClientResponse, error)
	Put(ctx context.Context, url string, body []byte, headers []Header) (ClientResponse, error)
}
