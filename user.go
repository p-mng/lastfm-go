package lastfmgo

// UserGetFriends calls the user.getFriends endpoint.
//
// https://www.last.fm/api/show/user.getFriends
func (c Client) UserGetFriends(params P) (UserGetFriendsResponse, error) {
	return getRequest[UserGetFriendsResponse](c.baseURL, params, "user.getFriends", c.apiKey, c.apiSecret, false)
}

// UserGetInfo calls the user.getInfo endpoint.
//
// https://www.last.fm/api/show/user.getInfo
func (c Client) UserGetInfo(params P) (UserGetInfoResponse, error) {
	return getRequest[UserGetInfoResponse](c.baseURL, params, "user.getInfo", c.apiKey, c.apiSecret, false)
}

// UserGetLovedTracks calls the user.getLovedTracks endpoint.
//
// https://www.last.fm/api/show/user.getLovedTracks
func (c Client) UserGetLovedTracks(params P) (UserGetLovedTracksResponse, error) {
	return getRequest[UserGetLovedTracksResponse](c.baseURL, params, "user.getLovedTracks", c.apiKey, c.apiSecret, false)
}

// UserGetPersonalTags calls the user.getPersonalTags endpoint.
//
// https://www.last.fm/api/show/user.getPersonalTags
func (c Client) UserGetPersonalTags(params P) (UserGetPersonalTagsResponse, error) {
	return getRequest[UserGetPersonalTagsResponse](c.baseURL, params, "user.getPersonalTags", c.apiKey, c.apiSecret, false)
}

// UserGetRecentTracks calls the user.getRecentTracks endpoint.
//
// https://www.last.fm/api/show/user.getRecentTracks
func (c Client) UserGetRecentTracks(params P) (UserGetRecentTracksResponse, error) {
	return getRequest[UserGetRecentTracksResponse](c.baseURL, params, "user.getRecentTracks", c.apiKey, c.apiSecret, false)
}

// UserGetTopAlbums calls the user.getTopAlbums endpoint.
//
// https://www.last.fm/api/show/user.getTopAlbums
func (c Client) UserGetTopAlbums(params P) (UserGetTopAlbumsResponse, error) {
	return getRequest[UserGetTopAlbumsResponse](c.baseURL, params, "user.getTopAlbums", c.apiKey, c.apiSecret, false)
}

// UserGetTopArtists calls the user.getTopArtists endpoint.
//
// https://www.last.fm/api/show/user.getTopArtists
func (c Client) UserGetTopArtists(params P) (UserGetTopArtistsResponse, error) {
	return getRequest[UserGetTopArtistsResponse](c.baseURL, params, "user.getTopArtists", c.apiKey, c.apiSecret, false)
}

// UserGetTopTags calls the user.getTopTags endpoint.
//
// https://www.last.fm/api/show/user.getTopTags
func (c Client) UserGetTopTags(params P) (UserGetTopTagsResponse, error) {
	return getRequest[UserGetTopTagsResponse](c.baseURL, params, "user.getTopTags", c.apiKey, c.apiSecret, false)
}

// UserGetTopTracks calls the user.getTopTracks endpoint.
//
// https://www.last.fm/api/show/user.getTopTracks
func (c Client) UserGetTopTracks(params P) (UserGetTopTracksResponse, error) {
	return getRequest[UserGetTopTracksResponse](c.baseURL, params, "user.getTopTracks", c.apiKey, c.apiSecret, false)
}

// UserGetWeeklyAlbumChart calls the user.getWeeklyAlbumChart endpoint.
//
// https://www.last.fm/api/show/user.getWeeklyAlbumChart
func (c Client) UserGetWeeklyAlbumChart(params P) (UserGetWeeklyAlbumChartResponse, error) {
	return getRequest[UserGetWeeklyAlbumChartResponse](c.baseURL, params, "user.getWeeklyAlbumChart", c.apiKey, c.apiSecret, false)
}

// UserGetWeeklyArtistChart calls the user.getWeeklyArtistChart endpoint.
//
// https://www.last.fm/api/show/user.getWeeklyArtistChart
func (c Client) UserGetWeeklyArtistChart(params P) (UserGetWeeklyArtistChartResponse, error) {
	return getRequest[UserGetWeeklyArtistChartResponse](c.baseURL, params, "user.getWeeklyArtistChart", c.apiKey, c.apiSecret, false)
}

// UserGetWeeklyChartList calls the user.getWeeklyChartList endpoint.
//
// https://www.last.fm/api/show/user.getWeeklyChartList
func (c Client) UserGetWeeklyChartList(params P) (UserGetWeeklyChartListResponse, error) {
	return getRequest[UserGetWeeklyChartListResponse](c.baseURL, params, "user.getWeeklyChartList", c.apiKey, c.apiSecret, false)
}

// UserGetWeeklyTrackChart calls the user.getWeeklyTrackChart endpoint.
//
// https://www.last.fm/api/show/user.getWeeklyTrackChart
func (c Client) UserGetWeeklyTrackChart(params P) (UserGetWeeklyTrackChartResponse, error) {
	return getRequest[UserGetWeeklyTrackChartResponse](c.baseURL, params, "user.getWeeklyTrackChart", c.apiKey, c.apiSecret, false)
}
