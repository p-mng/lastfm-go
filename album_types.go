package lastfmgo

// AlbumAddTagsResponse stores the response for album.addTags.
//
// https://www.last.fm/api/show/album.addTags
type AlbumAddTagsResponse struct {
	BaseResponse
}

// AlbumGetInfoResponse stores the response for album.getInfo.
//
// https://www.last.fm/api/show/album.getInfo
type AlbumGetInfoResponse struct {
	BaseResponse
	Album struct {
		Name   string `xml:"name"`
		Artist string `xml:"artist"`
		MBID   string `xml:"mbid"`
		URL    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Listeners     int64 `xml:"listeners"`
		Playcount     int64 `xml:"playcount"`
		UserPlaycount int64 `xml:"userplaycount"`
		Tracks        []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			URL        string `xml:"url"`
			Duration   int64  `xml:"duration"`
			Streamable struct {
				FullTrack  int64 `xml:"fulltrack,attr"`
				Streamable int64 `xml:",chardata"`
			} `xml:"streamable"`
			Artist struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
		} `xml:"tracks>track"`
		Tags []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"tags>tag"`
		Wiki struct {
			Published string `xml:"published"`
			Summary   string `xml:"summary"`
			Content   string `xml:"content"`
		} `xml:"wiki"`
	} `xml:"album"`
}

// AlbumGetTagsResponse stores the response for album.getTags.
//
// https://www.last.fm/api/show/album.getTags
type AlbumGetTagsResponse struct {
	BaseResponse
	Tags struct {
		Album  string `xml:"album,attr"`
		Artist string `xml:"artist,attr"`
		Tags   []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"tag"`
	} `xml:"tags"`
}

// AlbumGetTopTagsResponse stores the response for album.getTopTags.
//
// https://www.last.fm/api/show/album.getTopTags
type AlbumGetTopTagsResponse struct {
	BaseResponse
	TopTags struct {
		Album  string `xml:"album,attr"`
		Artist string `xml:"artist,attr"`
		Tags   []struct {
			Count int64  `xml:"count"`
			Name  string `xml:"name"`
			URL   string `xml:"url"`
		} `xml:"tag"`
	} `xml:"toptags"`
}

// AlbumRemoveTagResponse stores the response for album.removeTag.
//
// https://www.last.fm/api/show/album.removeTag
type AlbumRemoveTagResponse struct {
	BaseResponse
}

// AlbumSearchResponse stores the response for album.search.
//
// https://www.last.fm/api/show/album.search
type AlbumSearchResponse struct {
	BaseResponse
	Results struct {
		OpenSearchResponse
		Albums []struct {
			Name   string `xml:"name"`
			Artist string `xml:"artist"`
			URL    string `xml:"url"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			Streamable int64  `xml:"streamable"`
			MBID       string `xml:"mbid"`
		} `xml:"albummatches>album"`
	} `xml:"results"`
}
