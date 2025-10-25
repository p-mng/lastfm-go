package lastfmgo

// ChartGetTopArtistsResponse stores the response for chart.getTopArtists.
//
// https://www.last.fm/api/show/chart.getTopArtists
type ChartGetTopArtistsResponse struct {
	BaseResponse
	Artists struct {
		PaginationResponse
		Artists []struct {
			Name       string `xml:"name"`
			Playcount  int64  `xml:"playcount"`
			Listeners  int64  `xml:"listeners"`
			MBID       string `xml:"mbid"`
			URL        string `xml:"url"`
			Streamable int64  `xml:"streamable"`
			Images     []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"artist"`
	} `xml:"artists"`
}

// ChartGetTopTagsResponse stores the response for chart.getTopTags.
//
// https://www.last.fm/api/show/chart.getTopTags
type ChartGetTopTagsResponse struct {
	BaseResponse
	Tags struct {
		PaginationResponse
		Tags []struct {
			Name       string   `xml:"name"`
			URL        string   `xml:"url"`
			Reach      int64    `xml:"reach"`
			Taggings   int64    `xml:"taggings"`
			Streamable int64    `xml:"streamable"`
			Wiki       struct{} `xml:"wiki"`
		} `xml:"tag"`
	} `xml:"tags"`
}

// ChartGetTopTracksResponse stores the response for chart.getTopTracks.
//
// https://www.last.fm/api/show/chart.getTopTracks
type ChartGetTopTracksResponse struct {
	BaseResponse
	Tracks struct {
		PaginationResponse
		Tracks []struct {
			Name       string `xml:"name"`
			Duration   int64  `xml:"duration"`
			Playcount  int64  `xml:"playcount"`
			Listeners  int64  `xml:"listeners"`
			MBID       string `xml:"mbid"`
			URL        string `xml:"url"`
			Streamable struct {
				FullTrack  int64 `xml:"fulltrack,attr"`
				Streamable int64 `xml:",chardata"`
			} `xml:"streamable"`
			Artist struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"track"`
	} `xml:"tracks"`
}
