package lastfmgo

// AuthGetMobileSessionResponse stores the response for auth.getMobileSession.
//
// https://www.last.fm/api/show/auth.getMobileSession
type AuthGetMobileSessionResponse struct {
	BaseResponse
	Session struct {
		Name       string `xml:"name"`
		Key        string `xml:"key"`
		Subscriber int64  `xml:"subscriber"`
	} `xml:"session"`
}

// AuthGetSessionResponse stores the response for auth.getSession.
//
// https://www.last.fm/api/show/auth.getSession
type AuthGetSessionResponse struct {
	BaseResponse
	Session struct {
		Name       string `xml:"name"`
		Key        string `xml:"key"`
		Subscriber int64  `xml:"subscriber"`
	} `xml:"session"`
}

// AuthGetTokenResponse stores the response for auth.getToken.
//
// https://www.last.fm/api/show/auth.getToken
type AuthGetTokenResponse struct {
	BaseResponse
	Token string `xml:"token"`
}
