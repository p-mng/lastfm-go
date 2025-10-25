package lastfmgo

// TagGetInfoResponse stores the response for tag.getInfo.
//
// https://www.last.fm/api/show/tag.getInfo
type TagGetInfoResponse struct {
	BaseResponse
	Tag struct {
		Name  string `xml:"name"`
		Total int64  `xml:"total"`
		Reach int64  `xml:"reach"`
		Wiki  struct {
			Summary string `xml:"summary"`
			Content string `xml:"content"`
		} `xml:"wiki"`
	} `xml:"tag"`
}

// TagGetSimilarResponse stores the response for tag.getSimilar.
//
// https://www.last.fm/api/show/tag.getSimilar
type TagGetSimilarResponse struct {
	BaseResponse
	SimilarTags struct{} `xml:"similartags"`
}

// TagGetTopAlbumsResponse stores the response for tag.getTopAlbums.
//
// https://www.last.fm/api/show/tag.getTopAlbums
type TagGetTopAlbumsResponse struct {
	BaseResponse
	Albums struct {
		PaginationResponse
		Tag    string `xml:"tag,attr"`
		Albums []struct {
			Rank   int64  `xml:"rank,attr"`
			Name   string `xml:"name"`
			MBID   string `xml:"mbid"`
			URL    string `xml:"url"`
			Artist struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"album"`
	} `xml:"albums"`
}

// TagGetTopArtistsResponse stores the response for tag.getTopArtists.
//
// https://www.last.fm/api/show/tag.getTopArtists
type TagGetTopArtistsResponse struct {
	BaseResponse
	TopArtists struct {
		PaginationResponse
		Tag     string `xml:"tag,attr"`
		Artists []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
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

// TagGetTopTagsResponse stores the response for tag.getTopTags.
//
// https://www.last.fm/api/show/tag.getTopTags
type TagGetTopTagsResponse struct {
	BaseResponse
	TopTags struct {
		NumRes int64 `xml:"num_res,attr"`
		Offset int64 `xml:"offset,attr"`
		Total  int64 `xml:"total,attr"`
		Tags   []struct {
			Name  string `xml:"name"`
			Count int64  `xml:"count"`
			Reach int64  `xml:"reach"`
		} `xml:"tag"`
	} `xml:"toptags"`
}

// TagGetTopTracksResponse stores the response for tag.getTopTracks.
//
// https://www.last.fm/api/show/tag.getTopTracks
type TagGetTopTracksResponse struct {
	BaseResponse
	Tracks struct {
		PaginationResponse
		Tag    string `xml:"tag,attr"`
		Tracks []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Duration   int64  `xml:"duration"`
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
			MBID string `xml:"mbid"`
		} `xml:"track"`
	} `xml:"tracks"`
}

// TagGetWeeklyChartListResponse stores the response for tag.getWeeklyChartList.
//
// https://www.last.fm/api/show/tag.getWeeklyChartList
type TagGetWeeklyChartListResponse struct {
	BaseResponse
	Weeklychartlist struct {
		Tag    string `xml:"tag,attr"`
		Charts []struct {
			From int64 `xml:"from,attr"`
			To   int64 `xml:"to,attr"`
		} `xml:"chart"`
	} `xml:"weeklychartlist"`
}
