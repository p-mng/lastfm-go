package lastfmgo

// TrackAddTagsResponse stores the response for track.addTags.
//
// https://www.last.fm/api/show/track.addTags
type TrackAddTagsResponse struct {
	BaseResponse
}

// TrackGetCorrectionResponse stores the response for track.getCorrection.
//
// https://www.last.fm/api/show/track.getCorrection
type TrackGetCorrectionResponse struct {
	BaseResponse
	Correction struct {
		ArtistCorrected int64 `xml:"artistcorrected,attr"`
		Index           int64 `xml:"index,attr"`
		TrackCorrected  int64 `xml:"trackcorrected,attr"`
		Track           struct {
			Name   string `xml:"name"`
			URL    string `xml:"url"`
			Artist struct {
				Name string `xml:"name"`
				URL  string `xml:"url"`
			} `xml:"artist"`
		} `xml:"track"`
	} `xml:"corrections>correction"`
}

// TrackGetInfoResponse stores the response for track.getInfo.
//
// https://www.last.fm/api/show/track.getInfo
type TrackGetInfoResponse struct {
	BaseResponse
	Track struct {
		Name       string `xml:"name"`
		URL        string `xml:"url"`
		Duration   int64  `xml:"duration"`
		Streamable struct {
			FullTrack  int64 `xml:"fulltrack,attr"`
			Streamable int64 `xml:",chardata"`
		} `xml:"streamable"`
		Listeners int64 `xml:"listeners"`
		Playcount int64 `xml:"playcount"`
		Artist    struct {
			Name string `xml:"name"`
			MBID string `xml:"mbid"`
			URL  string `xml:"url"`
		} `xml:"artist"`
		Album struct {
			Artist string `xml:"artist"`
			Title  string `xml:"title"`
			URL    string `xml:"url"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"album"`
		Tags []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"toptags>tag"`
		Wiki struct {
			Published string `xml:"published"`
			Summary   string `xml:"summary"`
			Content   string `xml:"content"`
		} `xml:"wiki"`
	} `xml:"track"`
}

// TrackGetSimilarResponse stores the response for track.getSimilar.
//
// https://www.last.fm/api/show/track.getSimilar
type TrackGetSimilarResponse struct {
	BaseResponse
	SimilarTracks struct {
		Artist string `xml:"artist,attr"`
		Track  string `xml:"track,attr"`
		Tracks []struct {
			Name       string  `xml:"name"`
			Playcount  int64   `xml:"playcount"`
			Match      float64 `xml:"match"`
			URL        string  `xml:"url"`
			Streamable struct {
				FullTrack  int64 `xml:"fulltrack,attr"`
				Streamable int64 `xml:",chardata"`
			} `xml:"streamable"`
			Duration int64 `xml:"duration"`
			Artist   struct {
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
	} `xml:"similartracks"`
}

// TrackGetTagsResponse stores the response for track.getTags.
//
// https://www.last.fm/api/show/track.getTags
type TrackGetTagsResponse struct {
	BaseResponse
	Tags struct {
		Artist string `xml:"artist,attr"`
		Track  string `xml:"track,attr"`
		Tags   []struct {
			Name string `xml:"name"`
			URL  string `xml:"url"`
		} `xml:"tag"`
	} `xml:"tags"`
}

// TrackGetTopTagsResponse stores the response for track.getTopTags.
//
// https://www.last.fm/api/show/track.getTopTags
type TrackGetTopTagsResponse struct {
	BaseResponse
	TopTags struct {
		Artist string `xml:"artist,attr"`
		Track  string `xml:"track,attr"`
		Tags   []struct {
			Count int64  `xml:"count"`
			Name  string `xml:"name"`
			URL   string `xml:"url"`
		} `xml:"tag"`
	} `xml:"toptags"`
}

// TrackLoveResponse stores the response for track.love.
//
// https://www.last.fm/api/show/track.love
type TrackLoveResponse struct {
	BaseResponse
}

// TrackRemoveTagResponse stores the response for track.removeTag.
//
// https://www.last.fm/api/show/track.removeTag
type TrackRemoveTagResponse struct {
	BaseResponse
}

// TrackScrobbleResponse stores the response for track.scrobble.
//
// https://www.last.fm/api/show/track.scrobble
type TrackScrobbleResponse struct {
	BaseResponse
	Scrobbles struct {
		Accepted  int64 `xml:"accepted,attr"`
		Ignored   int64 `xml:"ignored,attr"`
		Scrobbles []struct {
			Track struct {
				Corrected int64  `xml:"corrected,attr"`
				Name      string `xml:",chardata"`
			} `xml:"track"`
			Artist struct {
				Corrected int64  `xml:"corrected,attr"`
				Name      string `xml:",chardata"`
			} `xml:"artist"`
			Album struct {
				Corrected int64  `xml:"corrected,attr"`
				Name      string `xml:",chardata"`
			} `xml:"album"`
			AlbumArtist struct {
				Corrected int64  `xml:"corrected,attr"`
				Name      string `xml:",chardata"`
			} `xml:"albumArtist"`
			Timestamp      int64 `xml:"timestamp"`
			IgnoredMessage struct {
				Code    int64  `xml:"code,attr"`
				Message string `xml:",chardata"`
			} `xml:"ignoredMessage"`
		} `xml:"scrobble"`
	} `xml:"scrobbles"`
}

// TrackSearchResponse stores the response for track.search.
//
// https://www.last.fm/api/show/track.search
type TrackSearchResponse struct {
	BaseResponse
	Results struct {
		OpenSearchResponse
		ItemsPerPage int64 `xml:"itemsPerPage"`
		Tracks       []struct {
			Name       string `xml:"name"`
			Artist     string `xml:"artist"`
			URL        string `xml:"url"`
			Streamable int64  `xml:"streamable"`
			Listeners  int64  `xml:"listeners"`
			Images     []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			MBID string `xml:"mbid"`
		} `xml:"trackmatches>track"`
	} `xml:"results"`
}

// TrackUnloveResponse stores the response for track.unlove.
//
// https://www.last.fm/api/show/track.unlove
type TrackUnloveResponse struct {
	BaseResponse
}

// TrackUpdateNowPlayingResponse stores the response for track.updateNowPlaying.
//
// https://www.last.fm/api/show/track.updateNowPlaying
type TrackUpdateNowPlayingResponse struct {
	BaseResponse
	Nowplaying struct {
		Track struct {
			Corrected int64  `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"track"`
		Artist struct {
			Corrected int64  `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"artist"`
		Album struct {
			Corrected int64 `xml:"corrected,attr"`
		} `xml:"album"`
		AlbumArtist struct {
			Corrected int64 `xml:"corrected,attr"`
		} `xml:"albumArtist"`
		IgnoredMessage struct {
			Code    int64  `xml:"code,attr"`
			Message string `xml:",chardata"`
		} `xml:"ignoredMessage"`
	} `xml:"nowplaying"`
}
