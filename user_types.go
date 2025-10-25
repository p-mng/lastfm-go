package lastfmgo

// UserGetFriendsResponse stores the response for user.getFriends.
//
// https://www.last.fm/api/show/user.getFriends
type UserGetFriendsResponse struct {
	BaseResponse
	Friends struct {
		PaginationResponse
		User  string `xml:"user,attr"`
		Users []struct {
			Name     string `xml:"name"`
			RealName string `xml:"realname"`
			Images   []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			URL        string `xml:"url"`
			Country    string `xml:"country"`
			Subscriber int64  `xml:"subscriber"`
			Playcount  int64  `xml:"playcount"`
			Playlists  int64  `xml:"playlists"`
			Bootstrap  int64  `xml:"bootstrap"`
			Registered struct {
				UnixTime int64  `xml:"unixtime,attr"`
				Time     string `xml:",chardata"`
			} `xml:"registered"`
			Type string `xml:"type"`
		} `xml:"user"`
	} `xml:"friends"`
}

// UserGetInfoResponse stores the response for user.getInfo.
//
// https://www.last.fm/api/show/user.getInfo
type UserGetInfoResponse struct {
	BaseResponse
	User struct {
		Name     string `xml:"name"`
		RealName string `xml:"realname"`
		Images   []struct {
			Size string `xml:"size,attr"`
			URL  string `xml:",chardata"`
		} `xml:"image"`
		URL        string `xml:"url"`
		Country    string `xml:"country"`
		Age        int64  `xml:"age"`
		Gender     string `xml:"gender"`
		Subscriber int64  `xml:"subscriber"`
		Playcount  int64  `xml:"playcount"`
		Playlists  int64  `xml:"playlists"`
		Bootstrap  int64  `xml:"bootstrap"`
		Registered struct {
			UnixTime int64  `xml:"unixtime,attr"`
			Time     string `xml:",chardata"`
		} `xml:"registered"`
		Type        string `xml:"type"`
		ArtistCount int64  `xml:"artist_count"`
		AlbumCount  int64  `xml:"album_count"`
		TrackCount  int64  `xml:"track_count"`
	} `xml:"user"`
}

// UserGetLovedTracksResponse stores the response for user.getLovedTracks.
//
// https://www.last.fm/api/show/user.getLovedTracks
type UserGetLovedTracksResponse struct {
	BaseResponse
	Lovedtracks struct {
		PaginationResponse
		User   string `xml:"user,attr"`
		Tracks []struct {
			Name string `xml:"name"`
			MBID string `xml:"mbid"`
			URL  string `xml:"url"`
			Date struct {
				UTS  int64  `xml:"uts,attr"`
				Time string `xml:",chardata"`
			} `xml:"date"`
			Artist struct {
				Name string `xml:"name"`
				MBID string `xml:"mbid"`
				URL  string `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			Streamable struct {
				FullTrack  int64 `xml:"fulltrack,attr"`
				Streamable int64 `xml:",chardata"`
			} `xml:"streamable"`
		} `xml:"track"`
	} `xml:"lovedtracks"`
}

// UserGetPersonalTagsResponse stores the response for user.getPersonalTags.
//
// https://www.last.fm/api/show/user.getPersonalTags
type UserGetPersonalTagsResponse struct {
	BaseResponse
	Taggings struct {
		PaginationResponse
		Tag    string `xml:"tag,attr"`
		User   string `xml:"user,attr"`
		Albums []struct {
			Name   string   `xml:"name"`
			MBID   struct{} `xml:"mbid"`
			URL    string   `xml:"url"`
			Artist struct {
				Name string   `xml:"name"`
				MBID struct{} `xml:"mbid"`
				URL  string   `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"albums>album"`
		Artists []struct {
			Name       string   `xml:"name"`
			MBID       struct{} `xml:"mbid"`
			URL        string   `xml:"url"`
			Streamable int64    `xml:"streamable"`
			Images     []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"artists>artist"`
		Tracks []struct {
			Name       string   `xml:"name"`
			Duration   string   `xml:"duration"`
			MBID       struct{} `xml:"mbid"`
			URL        string   `xml:"url"`
			Streamable struct {
				FullTrack  int64 `xml:"fulltrack,attr"`
				Streamable int64 `xml:",chardata"`
			} `xml:"streamable"`
			Artist struct {
				Name string   `xml:"name"`
				MBID struct{} `xml:"mbid"`
				URL  string   `xml:"url"`
			} `xml:"artist"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
		} `xml:"tracks>track"`
	} `xml:"taggings"`
}

// UserGetRecentTracksResponse stores the response for user.getRecentTracks.
//
// https://www.last.fm/api/show/user.getRecentTracks
type UserGetRecentTracksResponse struct {
	BaseResponse
	RecentTracks struct {
		PaginationResponse
		User   string `xml:"user,attr"`
		Tracks []struct {
			Artist struct {
				Name   string   `xml:"name"`
				MBID   struct{} `xml:"mbid"`
				URL    string   `xml:"url"`
				Images []struct {
					Size string `xml:"size,attr"`
					URL  string `xml:",chardata"`
				} `xml:"image"`
			} `xml:"artist"`
			Loved      int64  `xml:"loved"`
			Name       string `xml:"name"`
			Streamable int64  `xml:"streamable"`
			MBID       string `xml:"mbid"`
			Album      struct {
				MBID string `xml:"mbid,attr"`
				Name string `xml:",chardata"`
			} `xml:"album"`
			URL    string `xml:"url"`
			Images []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			Date struct {
				UTS  int64  `xml:"uts,attr"`
				Time string `xml:",chardata"`
			} `xml:"date"`
		} `xml:"track"`
	} `xml:"recenttracks"`
}

// UserGetTopAlbumsResponse stores the response for user.getTopAlbums.
//
// https://www.last.fm/api/show/user.getTopAlbums
type UserGetTopAlbumsResponse struct {
	BaseResponse
	TopAlbums struct {
		PaginationResponse
		User   string `xml:"user,attr"`
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

// UserGetTopArtistsResponse stores the response for user.getTopArtists.
//
// https://www.last.fm/api/show/user.getTopArtists
type UserGetTopArtistsResponse struct {
	BaseResponse
	TopArtists struct {
		PaginationResponse
		User    string `xml:"user,attr"`
		Artists []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Playcount  int64  `xml:"playcount"`
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

// UserGetTopTagsResponse stores the response for user.getTopTags.
//
// https://www.last.fm/api/show/user.getTopTags
type UserGetTopTagsResponse struct {
	BaseResponse
	TopTags struct {
		User string `xml:"user,attr"`
		Tags []struct {
			Name  string `xml:"name"`
			Count int64  `xml:"count"`
			URL   string `xml:"url"`
		} `xml:"tag"`
	} `xml:"toptags"`
}

// UserGetTopTracksResponse stores the response for user.getTopTracks.
//
// https://www.last.fm/api/show/user.getTopTracks
type UserGetTopTracksResponse struct {
	BaseResponse
	TopTracks struct {
		PaginationResponse
		User   string `xml:"user,attr"`
		Tracks []struct {
			Rank       int64  `xml:"rank,attr"`
			Name       string `xml:"name"`
			Duration   int64  `xml:"duration"`
			Playcount  int64  `xml:"playcount"`
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
	} `xml:"toptracks"`
}

// UserGetWeeklyAlbumChartResponse stores the response for user.getWeeklyAlbumChart.
//
// https://www.last.fm/api/show/user.getWeeklyAlbumChart
type UserGetWeeklyAlbumChartResponse struct {
	BaseResponse
	WeeklyAlbumChart struct {
		From   int64  `xml:"from,attr"`
		To     int64  `xml:"to,attr"`
		User   string `xml:"user,attr"`
		Albums []struct {
			Rank   int64 `xml:"rank,attr"`
			Artist struct {
				MBID string `xml:"mbid,attr"`
				Name string `xml:",chardata"`
			} `xml:"artist"`
			Name      string `xml:"name"`
			MBID      string `xml:"mbid"`
			Playcount int64  `xml:"playcount"`
			URL       string `xml:"url"`
		} `xml:"album"`
	} `xml:"weeklyalbumchart"`
}

// UserGetWeeklyArtistChartResponse stores the response for user.getWeeklyArtistChart.
//
// https://www.last.fm/api/show/user.getWeeklyArtistChart
type UserGetWeeklyArtistChartResponse struct {
	BaseResponse
	WeeklyArtistChart struct {
		From    int64  `xml:"from,attr"`
		To      int64  `xml:"to,attr"`
		User    string `xml:"user,attr"`
		Artists []struct {
			Rank      int64  `xml:"rank,attr"`
			Name      string `xml:"name"`
			MBID      string `xml:"mbid"`
			Playcount int64  `xml:"playcount"`
			URL       string `xml:"url"`
		} `xml:"artist"`
	} `xml:"weeklyartistchart"`
}

// UserGetWeeklyChartListResponse stores the response for user.getWeeklyChartList.
//
// https://www.last.fm/api/show/user.getWeeklyChartList
type UserGetWeeklyChartListResponse struct {
	BaseResponse
	Weeklychartlist struct {
		User   string `xml:"user,attr"`
		Charts []struct {
			From int64 `xml:"from,attr"`
			To   int64 `xml:"to,attr"`
		} `xml:"chart"`
	} `xml:"weeklychartlist"`
}

// UserGetWeeklyTrackChartResponse stores the response for user.getWeeklyTrackChart.
//
// https://www.last.fm/api/show/user.getWeeklyTrackChart
type UserGetWeeklyTrackChartResponse struct {
	BaseResponse
	WeeklyTrackChart struct {
		From   int64  `xml:"from,attr"`
		To     int64  `xml:"to,attr"`
		User   string `xml:"user,attr"`
		Tracks []struct {
			Rank   int64 `xml:"rank,attr"`
			Artist struct {
				MBID string `xml:"mbid,attr"`
				Name string `xml:",chardata"`
			} `xml:"artist"`
			Name      string `xml:"name"`
			MBID      string `xml:"mbid"`
			Playcount int64  `xml:"playcount"`
			Images    []struct {
				Size string `xml:"size,attr"`
				URL  string `xml:",chardata"`
			} `xml:"image"`
			URL string `xml:"url"`
		} `xml:"track"`
	} `xml:"weeklytrackchart"`
}
