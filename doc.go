// Package lastfmgo provides a Go wrapper for the last.fm REST API (2.0) and modern alternative to shkh/lastfm-go.
// This library aims to be idomatic, simple, and powerful at the same time.
//
// Generally, using this library involves creating a new client (web, mobile, or desktop).
// For write endpoints (e.g., track.scrobble or album.addTags), this client needs to be authenticated to generate a session key.
// This session key can be used alongside the other parameters to make authenticated requests.
// Please refer to the last.fm API documentation for the other required and optional parameters.
//
// The parameters "api_key" and "api_sig" are set automatically by the client and do not need to be specified.
// The client also automatically handles API errors and converts them to Go errors.
//
// Please refer to the examples for common use cases (e.g., authentication, scrobbling, error handling,
// and integration with other libraries).
package lastfmgo
