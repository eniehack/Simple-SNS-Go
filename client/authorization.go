// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Simple-SNS": Authorization Resource Client
//
// Command:
// $ goagen
// --design=github.com/eniehack/simple-sns-go/design
// --out=D:\project\simple-sns-go
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// RegisterAuthorizationPayload is the Authorization register action payload.
type RegisterAuthorizationPayload struct {
	Password   string `form:"password" json:"password" yaml:"password" xml:"password"`
	ScreenName string `form:"screen_name" json:"screen_name" yaml:"screen_name" xml:"screen_name"`
	Userid     string `form:"userid" json:"userid" yaml:"userid" xml:"userid"`
}

// RegisterAuthorizationPath computes a request path to the register action of Authorization.
func RegisterAuthorizationPath() string {

	return fmt.Sprintf("/api/v1/auth/new")
}

// 新規登録
func (c *Client) RegisterAuthorization(ctx context.Context, path string, payload *RegisterAuthorizationPayload) (*http.Response, error) {
	req, err := c.NewRegisterAuthorizationRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterAuthorizationRequest create the request corresponding to the register action endpoint of the Authorization resource.
func (c *Client) NewRegisterAuthorizationRequest(ctx context.Context, path string, payload *RegisterAuthorizationPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}