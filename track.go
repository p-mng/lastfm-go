package lastfmgo

// TrackAddTags calls the track.addTags endpoint.
//
// https://www.last.fm/api/show/track.addTags
func (c Client) TrackAddTags(params P) (TrackAddTagsResponse, error) {
	return PostRequest[TrackAddTagsResponse](c.baseURL, params, "track.addTags", c.apiKey, c.apiSecret, true)
}

// TrackGetCorrection calls the track.getCorrection endpoint.
//
// https://www.last.fm/api/show/track.getCorrection
func (c Client) TrackGetCorrection(params P) (TrackGetCorrectionResponse, error) {
	return GetRequest[TrackGetCorrectionResponse](c.baseURL, params, "track.getCorrection", c.apiKey, c.apiSecret, false)
}

// TrackGetInfo calls the track.getInfo endpoint.
//
// https://www.last.fm/api/show/track.getInfo
func (c Client) TrackGetInfo(params P) (TrackGetInfoResponse, error) {
	return GetRequest[TrackGetInfoResponse](c.baseURL, params, "track.getInfo", c.apiKey, c.apiSecret, false)
}

// TrackGetSimilar calls the track.getSimilar endpoint.
//
// https://www.last.fm/api/show/track.getSimilar
func (c Client) TrackGetSimilar(params P) (TrackGetSimilarResponse, error) {
	return GetRequest[TrackGetSimilarResponse](c.baseURL, params, "track.getSimilar", c.apiKey, c.apiSecret, false)
}

// TrackGetTags calls the track.getTags endpoint.
//
// https://www.last.fm/api/show/track.getTags
func (c Client) TrackGetTags(params P) (TrackGetTagsResponse, error) {
	return GetRequest[TrackGetTagsResponse](c.baseURL, params, "track.getTags", c.apiKey, c.apiSecret, false)
}

// TrackGetTopTags calls the track.getTopTags endpoint.
//
// https://www.last.fm/api/show/track.getTopTags
func (c Client) TrackGetTopTags(params P) (TrackGetTopTagsResponse, error) {
	return GetRequest[TrackGetTopTagsResponse](c.baseURL, params, "track.getTopTags", c.apiKey, c.apiSecret, false)
}

// TrackLove calls the track.love endpoint.
//
// https://www.last.fm/api/show/track.love
func (c Client) TrackLove(params P) (TrackLoveResponse, error) {
	return PostRequest[TrackLoveResponse](c.baseURL, params, "track.love", c.apiKey, c.apiSecret, true)
}

// TrackRemoveTag calls the track.removeTag endpoint.
//
// https://www.last.fm/api/show/track.removeTag
func (c Client) TrackRemoveTag(params P) (TrackRemoveTagResponse, error) {
	return PostRequest[TrackRemoveTagResponse](c.baseURL, params, "track.removeTag", c.apiKey, c.apiSecret, true)
}

// TrackScrobble calls the track.scrobble endpoint.
//
// https://www.last.fm/api/show/track.scrobble
func (c Client) TrackScrobble(params P) (TrackScrobbleResponse, error) {
	return PostRequest[TrackScrobbleResponse](c.baseURL, params, "track.scrobble", c.apiKey, c.apiSecret, true)
}

// TrackSearch calls the track.search endpoint.
//
// https://www.last.fm/api/show/track.search
func (c Client) TrackSearch(params P) (TrackSearchResponse, error) {
	return GetRequest[TrackSearchResponse](c.baseURL, params, "track.search", c.apiKey, c.apiSecret, false)
}

// TrackUnlove calls the track.unlove endpoint.
//
// https://www.last.fm/api/show/track.unlove
func (c Client) TrackUnlove(params P) (TrackUnloveResponse, error) {
	return PostRequest[TrackUnloveResponse](c.baseURL, params, "track.unlove", c.apiKey, c.apiSecret, true)
}

// TrackUpdateNowPlaying calls the track.updateNowPlaying endpoint.
//
// https://www.last.fm/api/show/track.updateNowPlaying
func (c Client) TrackUpdateNowPlaying(params P) (TrackUpdateNowPlayingResponse, error) {
	return PostRequest[TrackUpdateNowPlayingResponse](c.baseURL, params, "track.updateNowPlaying", c.apiKey, c.apiSecret, true)
}
