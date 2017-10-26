// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: login Resource Client
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/authorization-server/design
// --out=$(GOPATH)/src/github.com/JormungandrK/authorization-server
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// ShowLoginLoginPath computes a request path to the showLogin action of login.
func ShowLoginLoginPath() string {

	return fmt.Sprintf("/auth/login")
}

// Shows a login screen
func (c *Client) ShowLoginLogin(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowLoginLoginRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowLoginLoginRequest create the request corresponding to the showLogin action endpoint of the login resource.
func (c *Client) NewShowLoginLoginRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
