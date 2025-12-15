package lastfmgo

// GeoGetTopArtists calls the geo.getTopArtists endpoint.
//
// https://www.last.fm/api/show/geo.getTopArtists
func (c Client) GeoGetTopArtists(params P) (GeoGetTopArtistsResponse, error) {
	return GetRequest[GeoGetTopArtistsResponse](c.baseURL, params, "geo.getTopArtists", c.apiKey, c.apiSecret, false)
}

// GeoGetTopTracks calls the geo.getTopTracks endpoint.
//
// https://www.last.fm/api/show/geo.getTopTracks
func (c Client) GeoGetTopTracks(params P) (GeoGetTopTracksResponse, error) {
	return GetRequest[GeoGetTopTracksResponse](c.baseURL, params, "geo.getTopTracks", c.apiKey, c.apiSecret, false)
}
