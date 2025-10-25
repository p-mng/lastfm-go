package lastfmgo

// LibraryGetArtistsResponse stores the response for library.getArtists.
//
// https://www.last.fm/api/show/library.getArtists
type LibraryGetArtistsResponse struct {
	BaseResponse
	Artists struct {
		PaginationResponse
		User    string `xml:"user,attr"`
		Artists []struct {
			Name       string `xml:"name"`
			Playcount  int64  `xml:"playcount"`
			Tagcount   int64  `xml:"tagcount"`
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
