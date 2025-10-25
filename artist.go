package lastfmgo

// ArtistAddTags calls the artist.addTags endpoint.
//
// https://www.last.fm/api/show/artist.addTags
func (c Client) ArtistAddTags(params P) (ArtistAddTagsResponse, error) {
	return postRequest[ArtistAddTagsResponse](c.baseURL, params, "artist.addTags", c.apiKey, c.apiSecret, true)
}

// ArtistGetCorrection calls the artist.getCorrection endpoint.
//
// https://www.last.fm/api/show/artist.getCorrection
func (c Client) ArtistGetCorrection(params P) (ArtistGetCorrectionResponse, error) {
	return getRequest[ArtistGetCorrectionResponse](c.baseURL, params, "artist.getCorrection", c.apiKey, c.apiSecret, false)
}

// ArtistGetInfo calls the artist.getInfo endpoint.
//
// https://www.last.fm/api/show/artist.getInfo
func (c Client) ArtistGetInfo(params P) (ArtistGetInfoResponse, error) {
	return getRequest[ArtistGetInfoResponse](c.baseURL, params, "artist.getInfo", c.apiKey, c.apiSecret, false)
}

// ArtistGetSimilar calls the artist.getSimilar endpoint.
//
// https://www.last.fm/api/show/artist.getSimilar
func (c Client) ArtistGetSimilar(params P) (ArtistGetSimilarResponse, error) {
	return getRequest[ArtistGetSimilarResponse](c.baseURL, params, "artist.getSimilar", c.apiKey, c.apiSecret, false)
}

// ArtistGetTags calls the artist.getTags endpoint.
//
// https://www.last.fm/api/show/artist.getTags
func (c Client) ArtistGetTags(params P) (ArtistGetTagsResponse, error) {
	return getRequest[ArtistGetTagsResponse](c.baseURL, params, "artist.getTags", c.apiKey, c.apiSecret, false)
}

// ArtistGetTopAlbums calls the artist.getTopAlbums endpoint.
//
// https://www.last.fm/api/show/artist.getTopAlbums
func (c Client) ArtistGetTopAlbums(params P) (ArtistGetTopAlbumsResponse, error) {
	return getRequest[ArtistGetTopAlbumsResponse](c.baseURL, params, "artist.getTopAlbums", c.apiKey, c.apiSecret, false)
}

// ArtistGetTopTags calls the artist.getTopTags endpoint.
//
// https://www.last.fm/api/show/artist.getTopTags
func (c Client) ArtistGetTopTags(params P) (ArtistGetTopTagsResponse, error) {
	return getRequest[ArtistGetTopTagsResponse](c.baseURL, params, "artist.getTopTags", c.apiKey, c.apiSecret, false)
}

// ArtistGetTopTracks calls the artist.getTopTracks endpoint.
//
// https://www.last.fm/api/show/artist.getTopTracks
func (c Client) ArtistGetTopTracks(params P) (ArtistGetTopTracksResponse, error) {
	return getRequest[ArtistGetTopTracksResponse](c.baseURL, params, "artist.getTopTracks", c.apiKey, c.apiSecret, false)
}

// ArtistRemoveTag calls the artist.removeTag endpoint.
//
// https://www.last.fm/api/show/artist.removeTag
func (c Client) ArtistRemoveTag(params P) (ArtistRemoveTagResponse, error) {
	return postRequest[ArtistRemoveTagResponse](c.baseURL, params, "artist.removeTag", c.apiKey, c.apiSecret, true)
}

// ArtistSearch calls the artist.search endpoint.
//
// https://www.last.fm/api/show/artist.search
func (c Client) ArtistSearch(params P) (ArtistSearchResponse, error) {
	return getRequest[ArtistSearchResponse](c.baseURL, params, "artist.search", c.apiKey, c.apiSecret, false)
}
