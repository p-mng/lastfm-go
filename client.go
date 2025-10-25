package lastfmgo

import (
	"errors"

	"github.com/p-mng/lastfm-go/internal"
)

// Client is a client for the last.fm API.
type Client struct {
	baseURL     string
	apiKey      string
	apiSecret   string
	username    string
	password    string
	callbackURL string
}

// NewWebClient creates a new client prepared for web authentication. Requires calling the following methods:
//
//  1. [Client.WebAuthorizationURL]
//  2. [Client.AuthGetSession]
//
// You must create a server running on callbackURL to catch the token.
// If callbackURL is an empty string, the default callback URL will be used.
//
// https://www.last.fm/api/webauth
func NewWebClient(baseURL, apiKey, apiSecret, callbackURL string) (Client, error) {
	switch {
	case !internal.IsURL(baseURL):
		return Client{}, errors.New("invalid URL")
	case !internal.IsMD5(apiKey):
		return Client{}, errors.New("invalid API key")
	case !internal.IsMD5(apiSecret):
		return Client{}, errors.New("invalid API secret")
	case !internal.IsURL(callbackURL):
		return Client{}, errors.New("invalid URL")
	}

	return Client{
		baseURL:     baseURL,
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		username:    "",
		password:    "",
		callbackURL: callbackURL,
	}, nil
}

// NewMobileClient creates a new client prepared for mobile authentication. Requires calling the following methods:
//
//  1. [Client.AuthGetMobileSession]
//
// https://www.last.fm/api/mobileauth
func NewMobileClient(baseURL, apiKey, apiSecret, username, password string) (Client, error) {
	switch {
	case !internal.IsURL(baseURL):
		return Client{}, errors.New("invalid URL")
	case !internal.IsMD5(apiKey):
		return Client{}, errors.New("invalid API key")
	case !internal.IsMD5(apiSecret):
		return Client{}, errors.New("invalid API secret")
	case username == "":
		return Client{}, errors.New("invalid username")
	case password == "":
		return Client{}, errors.New("invalid password")
	}

	return Client{
		baseURL:     baseURL,
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		username:    username,
		password:    password,
		callbackURL: "",
	}, nil
}

// NewDesktopClient creates a new client prepared for desktop authentication. Requires calling the following methods:
//
//  1. [Client.AuthGetToken]
//  2. [Client.DesktopAuthorizationURL]
//  3. [Client.AuthGetSession]
//
// https://www.last.fm/api/desktopauth
func NewDesktopClient(baseURL, apiKey, apiSecret string) (Client, error) {
	switch {
	case !internal.IsURL(baseURL):
		return Client{}, errors.New("invalid URL")
	case !internal.IsMD5(apiKey):
		return Client{}, errors.New("invalid API key")
	case !internal.IsMD5(apiSecret):
		return Client{}, errors.New("invalid API secret")
	}

	return Client{
		baseURL:     baseURL,
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		username:    "",
		password:    "",
		callbackURL: "",
	}, nil
}
