package models

import "net/http"

type ClientOpts struct {
	APIKey     string
	HTTPClient *http.Client
}

type Client struct {
	hc   *http.Client
	opts *ClientOpts
}
