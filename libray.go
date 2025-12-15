package lastfmgo

// LibraryGetArtists calls the library.getArtists endpoint.
//
// https://www.last.fm/api/show/library.getArtists
func (c Client) LibraryGetArtists(params P) (LibraryGetArtistsResponse, error) {
	return GetRequest[LibraryGetArtistsResponse](c.baseURL, params, "library.getArtists", c.apiKey, c.apiSecret, false)
}
