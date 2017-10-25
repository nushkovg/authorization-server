// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// unnamed API: authUI Resource Client
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
	"strconv"
)

// ConfirmAuthorizationAuthUIPath computes a request path to the confirmAuthorization action of authUI.
func ConfirmAuthorizationAuthUIPath() string {

	return fmt.Sprintf("/confirm-authorization")
}

// Confirm the authorization of the client
func (c *Client) ConfirmAuthorizationAuthUI(ctx context.Context, path string, confirmed *bool) (*http.Response, error) {
	req, err := c.NewConfirmAuthorizationAuthUIRequest(ctx, path, confirmed)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewConfirmAuthorizationAuthUIRequest create the request corresponding to the confirmAuthorization action endpoint of the authUI resource.
func (c *Client) NewConfirmAuthorizationAuthUIRequest(ctx context.Context, path string, confirmed *bool) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if confirmed != nil {
		tmp7 := strconv.FormatBool(*confirmed)
		values.Set("confirmed", tmp7)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PromptAuthorizationAuthUIPath computes a request path to the promptAuthorization action of authUI.
func PromptAuthorizationAuthUIPath() string {

	return fmt.Sprintf("/authorize-client")
}

// Prompt the user for client authorization
func (c *Client) PromptAuthorizationAuthUI(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPromptAuthorizationAuthUIRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPromptAuthorizationAuthUIRequest create the request corresponding to the promptAuthorization action endpoint of the authUI resource.
func (c *Client) NewPromptAuthorizationAuthUIRequest(ctx context.Context, path string) (*http.Request, error) {
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
