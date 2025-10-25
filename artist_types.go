package lastfmgo

// ArtistAddTagsResponse stores the response for artist.addTags.
//
// https://www.last.fm/api/show/artist.addTags
type ArtistAddTagsResponse struct {
	BaseResponse
}

// ArtistGetCorrectionResponse stores the response for artist.getCorrection.
//
// https://www.last.fm/api/show/artist.getCorrection
type ArtistGetCorrectionResponse struct {
	BaseResponse
	Correction struct {
		Index  int64 `xml:"index,attr"`
		Artist struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"artist"`
	} `xml:"corrections>correction"`
}

// ArtistGetInfoResponse stores the response for artist.getInfo.
//
// https://www.last.fm/api/show/artist.getInfo
type ArtistGetInfoResponse struct {
	BaseResponse
	Artist struct {
		Name   string `xml:"name"`
		MBID   string `xml:"mbid"`
		URL    string `xml:"url"`
		Images []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		Streamable int64 `xml:"streamable"`
		Ontour     int64 `xml:"ontour"`
		Stats      struct {
			Listeners     int64 `xml:"listeners"`
			Playcount     int64 `xml:"playcount"`
			UserPlaycount int64 `xml:"userplaycount"`
		} `xml:"stats"`
		Artists []struct {
			Name   string `xml:"name"`
			URL    string `xml:"url"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"similar>artist"`
		Tags []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"tags>tag"`
		Bio struct {
			Link struct {
				Href string `xml:"href,attr"`
				Rel  string `xml:"rel,attr"`
			} `xml:"links>link"`
			Published string `xml:"published"`
			Summary   string `xml:"summary"`
			Content   string `xml:"content"`
		} `xml:"bio"`
	} `xml:"artist"`
}

// ArtistGetSimilarResponse stores the response for artist.getSimilar.
//
// https://www.last.fm/api/show/artist.getSimilar
type ArtistGetSimilarResponse struct {
	BaseResponse
	SimilarArtists struct {
		Artist  string `xml:"artist,attr"`
		Artists []struct {
			Name   string  `xml:"name"`
			MBID   string  `xml:"mbid"`
			Match  float64 `xml:"match"`
			URL    string  `xml:"url"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			Streamable int64 `xml:"streamable"`
		} `xml:"artist"`
	} `xml:"similarartists"`
}

// ArtistGetTagsResponse stores the response for artist.getTags.
//
// https://www.last.fm/api/show/artist.getTags
type ArtistGetTagsResponse struct {
	BaseResponse
	Tags struct {
		Artist string `xml:"artist,attr"`
		Tags   []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"tag"`
	} `xml:"tags"`
}

// ArtistGetTopAlbumsResponse stores the response for artist.getTopAlbums.
//
// https://www.last.fm/api/show/artist.getTopAlbums
type ArtistGetTopAlbumsResponse struct {
	BaseResponse
	TopAlbums struct {
		PaginationResponse
		Artist string `xml:"artist,attr"`
		Albums []struct {
			Rank      int64  `xml:"rank,attr"`
			Name      string `xml:"name"`
			Playcount int64  `xml:"playcount"`
			MBID      string `xml:"mbid"`
			URL       string `xml:"url"`
			Artist    struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"album"`
	} `xml:"topalbums"`
}

// ArtistGetTopTagsResponse stores the response for artist.getTopTags.
//
// https://www.last.fm/api/show/artist.getTopTags
type ArtistGetTopTagsResponse struct {
	BaseResponse
	TopTags struct {
		Artist string `xml:"artist,attr"`
		Tags   []struct {
			Count int64  `xml:"count"`
			Name  string `xml:"name"`
			URL   string `xml:"url"`
		} `xml:"tag"`
	} `xml:"toptags"`
}

// ArtistGetTopTracksResponse stores the response for artist.getTopTracks.
//
// https://www.last.fm/api/show/artist.getTopTracks
type ArtistGetTopTracksResponse struct {
	BaseResponse
	TopTracks struct {
		PaginationResponse
		Artist string `xml:"artist,attr"`
		Tracks []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Playcount  int64  `xml:"playcount"`
			Listeners  int64  `xml:"listeners"`
			MBID       string `xml:"mbid"`
			URL        string `xml:"url"`
			Streamable int64  `xml:"streamable"`
			Artist     struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"track"`
	} `xml:"toptracks"`
}

// ArtistRemoveTagResponse stores the response for artist.removeTag.
//
// https://www.last.fm/api/show/artist.removeTag
type ArtistRemoveTagResponse struct {
	BaseResponse
}

// ArtistSearchResponse stores the response for artist.search.
//
// https://www.last.fm/api/show/artist.search
type ArtistSearchResponse struct {
	BaseResponse
	Results struct {
		OpenSearchResponse
		Artists []struct {
			Name       string `xml:"name"`
			Listeners  int64  `xml:"listeners"`
			MBID       string `xml:"mbid"`
			URL        string `xml:"url"`
			Streamable int64  `xml:"streamable"`
			Images     []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"artistmatches>artist"`
	} `xml:"results"`
}
