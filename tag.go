package lastfmgo

// TagGetInfo calls the tag.getInfo endpoint.
//
// https://www.last.fm/api/show/tag.getInfo
func (c Client) TagGetInfo(params P) (TagGetInfoResponse, error) {
	return getRequest[TagGetInfoResponse](c.baseURL, params, "tag.getInfo", c.apiKey, c.apiSecret, false)
}

// TagGetSimilar calls the tag.getSimilar endpoint.
//
// https://www.last.fm/api/show/tag.getSimilar
func (c Client) TagGetSimilar(params P) (TagGetSimilarResponse, error) {
	return getRequest[TagGetSimilarResponse](c.baseURL, params, "tag.getSimilar", c.apiKey, c.apiSecret, false)
}

// TagGetTopAlbums calls the tag.getTopAlbums endpoint.
//
// https://www.last.fm/api/show/tag.getTopAlbums
func (c Client) TagGetTopAlbums(params P) (TagGetTopAlbumsResponse, error) {
	return getRequest[TagGetTopAlbumsResponse](c.baseURL, params, "tag.getTopAlbums", c.apiKey, c.apiSecret, false)
}

// TagGetTopArtists calls the tag.getTopArtists endpoint.
//
// https://www.last.fm/api/show/tag.getTopArtists
func (c Client) TagGetTopArtists(params P) (TagGetTopArtistsResponse, error) {
	return getRequest[TagGetTopArtistsResponse](c.baseURL, params, "tag.getTopArtists", c.apiKey, c.apiSecret, false)
}

// TagGetTopTags calls the tag.getTopTags endpoint.
//
// https://www.last.fm/api/show/tag.getTopTags
func (c Client) TagGetTopTags(params P) (TagGetTopTagsResponse, error) {
	return getRequest[TagGetTopTagsResponse](c.baseURL, params, "tag.getTopTags", c.apiKey, c.apiSecret, false)
}

// TagGetTopTracks calls the tag.getTopTracks endpoint.
//
// https://www.last.fm/api/show/tag.getTopTracks
func (c Client) TagGetTopTracks(params P) (TagGetTopTracksResponse, error) {
	return getRequest[TagGetTopTracksResponse](c.baseURL, params, "tag.getTopTracks", c.apiKey, c.apiSecret, false)
}

// TagGetWeeklyChartList calls the tag.getWeeklyChartList endpoint.
//
// https://www.last.fm/api/show/tag.getWeeklyChartList
func (c Client) TagGetWeeklyChartList(params P) (TagGetWeeklyChartListResponse, error) {
	return getRequest[TagGetWeeklyChartListResponse](c.baseURL, params, "tag.getWeeklyChartList", c.apiKey, c.apiSecret, false)
}
