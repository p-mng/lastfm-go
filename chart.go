package lastfmgo

// ChartGetTopArtists calls the chart.getTopArtists endpoint.
//
// https://www.last.fm/api/show/chart.getTopArtists
func (c Client) ChartGetTopArtists(params P) (ChartGetTopArtistsResponse, error) {
	return GetRequest[ChartGetTopArtistsResponse](c.baseURL, params, "chart.getTopArtists", c.apiKey, c.apiSecret, false)
}

// ChartGetTopTags calls the chart.getTopTags endpoint.
//
// https://www.last.fm/api/show/chart.getTopTags
func (c Client) ChartGetTopTags(params P) (ChartGetTopTagsResponse, error) {
	return GetRequest[ChartGetTopTagsResponse](c.baseURL, params, "chart.getTopTags", c.apiKey, c.apiSecret, false)
}

// ChartGetTopTracks calls the chart.getTopTracks endpoint.
//
// https://www.last.fm/api/show/chart.getTopTracks
func (c Client) ChartGetTopTracks(params P) (ChartGetTopTracksResponse, error) {
	return GetRequest[ChartGetTopTracksResponse](c.baseURL, params, "chart.getTopTracks", c.apiKey, c.apiSecret, false)
}
