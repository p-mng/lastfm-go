package lastfmgo

// GeoGetTopArtistsResponse stores the response for geo.getTopArtists.
//
// https://www.last.fm/api/show/geo.getTopArtists
type GeoGetTopArtistsResponse struct {
	BaseResponse
	TopArtists struct {
		PaginationResponse
		Country string `xml:"country,attr"`
		Artists []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Listeners  int64  `xml:"listeners"`
			MBID       string `xml:"mbid"`
			URL        string `xml:"url"`
			Streamable int64  `xml:"streamable"`
			Images     []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"artist"`
	} `xml:"topartists"`
}

// GeoGetTopTracksResponse stores the response for geo.getTopTracks.
//
// https://www.last.fm/api/show/geo.getTopTracks
type GeoGetTopTracksResponse struct {
	BaseResponse
	Tracks struct {
		PaginationResponse
		Country string `xml:"country,attr"`
		Tracks  []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Duration   int64  `xml:"duration"`
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
