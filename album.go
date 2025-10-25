package lastfmgo

// AlbumAddTags calls the album.addTags endpoint.
//
// https://www.last.fm/api/show/album.addTags
func (c Client) AlbumAddTags(params P) (AlbumAddTagsResponse, error) {
	return postRequest[AlbumAddTagsResponse](c.baseURL, params, "album.addTags", c.apiKey, c.apiSecret, true)
}

// AlbumGetInfo calls the album.getInfo endpoint.
//
// https://www.last.fm/api/show/album.getInfo
func (c Client) AlbumGetInfo(params P) (AlbumGetInfoResponse, error) {
	return getRequest[AlbumGetInfoResponse](c.baseURL, params, "album.getInfo", c.apiKey, c.apiSecret, false)
}

// AlbumGetTags calls the album.getTags endpoint.
//
// https://www.last.fm/api/show/album.getTags
func (c Client) AlbumGetTags(params P) (AlbumGetTagsResponse, error) {
	return getRequest[AlbumGetTagsResponse](c.baseURL, params, "album.getTags", c.apiKey, c.apiSecret, false)
}

// AlbumGetTopTags calls the album.getTopTags endpoint.
//
// https://www.last.fm/api/show/album.getTopTags
func (c Client) AlbumGetTopTags(params P) (AlbumGetTopTagsResponse, error) {
	return getRequest[AlbumGetTopTagsResponse](c.baseURL, params, "album.getTopTags", c.apiKey, c.apiSecret, false)
}

// AlbumRemoveTag calls the album.removeTag endpoint.
//
// https://www.last.fm/api/show/album.removeTag
func (c Client) AlbumRemoveTag(params P) (AlbumRemoveTagResponse, error) {
	return postRequest[AlbumRemoveTagResponse](c.baseURL, params, "album.removeTag", c.apiKey, c.apiSecret, true)
}

// AlbumSearch calls the album.search endpoint.
//
// https://www.last.fm/api/show/album.search
func (c Client) AlbumSearch(params P) (AlbumSearchResponse, error) {
	return getRequest[AlbumSearchResponse](c.baseURL, params, "album.search", c.apiKey, c.apiSecret, false)
}
