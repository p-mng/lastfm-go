package lastfmgo

import (
	"fmt"
)

// AuthGetToken calls the auth.getToken endpoint.
// This is the first step in the desktop authentication process.
//
// https://www.last.fm/api/show/auth.getToken
func (c Client) AuthGetToken() (AuthGetTokenResponse, error) {
	return GetRequest[AuthGetTokenResponse](c.baseURL, P{}, "auth.getToken", c.apiKey, c.apiSecret, false)
}

// AuthGetSession calls the auth.getSession endpoint.
// This is the second step in the web authentication process and the third step in the desktop authentication process.
//
// https://www.last.fm/api/show/auth.getSession
func (c Client) AuthGetSession(token string) (AuthGetSessionResponse, error) {
	return GetRequest[AuthGetSessionResponse](c.baseURL, P{
		"token": token,
	}, "auth.getSession", c.apiKey, c.apiSecret, true)
}

// AuthGetMobileSession calls the auth.getMobileSession endpoint.
// This is the first and only step in the mobile authentication process.
//
// https://www.last.fm/api/show/auth.getMobileSession
func (c Client) AuthGetMobileSession() (AuthGetMobileSessionResponse, error) {
	return PostRequest[AuthGetMobileSessionResponse](c.baseURL, P{
		"password": c.password,
		"username": c.username,
	}, "auth.getMobileSession", c.apiKey, c.apiSecret, true)
}

// DesktopAuthorizationURL generates a URL to request authorization from the user.
// This is the second step in the desktop authentication process.
//
// https://www.last.fm/api/desktopauth
func (c Client) DesktopAuthorizationURL(token string) string {
	return fmt.Sprintf("http://www.last.fm/api/auth/?api_key=%s&token=%s", c.apiKey, token)
}

// WebAuthorizationURL generates a URL to request authorization from the user.
// This is the first step in the web authentication process.
//
// https://www.last.fm/api/webauth
func (c Client) WebAuthorizationURL() string {
	if c.callbackURL == "" {
		return fmt.Sprintf("http://www.last.fm/api/auth/?api_key=%s", c.apiKey)
	}
	return fmt.Sprintf("http://www.last.fm/api/auth/?api_key=%s&cb=%s", c.apiKey, c.callbackURL)
}
