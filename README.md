# lastfm-go

A Go wrapper for the last.fm REST API (2.0) and modern alternative to [shkh/lastfm-go](https://github.com/shkh/lastfm-go).
This library aims to be idomatic, simple, and powerful at the same time.

> [!NOTE]
> While the `go.mod` file contains a number of dependencies for testing and examples, the core library itself has no external dependencies.

## Installation

```shell
go get -u github.com/p-mng/lastfm-go
```

## Usage

Please refer to the provided examples for common use cases (e.g., authentication, scrobbling, and integration with other libraries).

## Supported API endpoints

### `album`

- [x] `album.addTags`
- [x] `album.getInfo`
- [x] `album.getTags`
- [x] `album.getTopTags`
- [x] `album.removeTag`
- [x] `album.search`

### `artist`

- [x] `artist.addTags`
- [x] `artist.getCorrection`
- [x] `artist.getInfo`
- [x] `artist.getSimilar`
- [x] `artist.getTags`
- [x] `artist.getTopAlbums`
- [x] `artist.getTopTags`
- [x] `artist.getTopTracks`
- [x] `artist.removeTag`
- [x] `artist.search`

### `auth`

- [x] `auth.getMobileSession`
- [x] `auth.getSession`
- [x] `auth.getToken`

### `chart`

- [x] `chart.getTopArtists`
- [x] `chart.getTopTags`
- [x] `chart.getTopTracks`

### `geo`

- [x] `geo.getTopArtists`
- [x] `geo.getTopTracks`

### `library`

- [x] `library.getArtists`

### `tag`

- [x] `tag.getInfo`
- [x] `tag.getSimilar`
- [x] `tag.getTopAlbums`
- [x] `tag.getTopArtists`
- [x] `tag.getTopTags`
- [x] `tag.getTopTracks`
- [x] `tag.getWeeklyChartList`

### `track`

- [x] `track.addTags`
- [x] `track.getCorrection`
- [x] `track.getInfo`
- [x] `track.getSimilar`
- [x] `track.getTags`
- [x] `track.getTopTags`
- [x] `track.love`
- [x] `track.removeTag`
- [x] `track.scrobble`
- [x] `track.search`
- [x] `track.unlove`
- [x] `track.updateNowPlaying`

### `user`

- [x] `user.getFriends`
- [x] `user.getInfo`
- [x] `user.getLovedTracks`
- [x] `user.getPersonalTags`
- [x] `user.getRecentTracks`
- [x] `user.getTopAlbums`
- [x] `user.getTopArtists`
- [x] `user.getTopTags`
- [x] `user.getTopTracks`
- [x] `user.getWeeklyAlbumChart`
- [x] `user.getWeeklyArtistChart`
- [x] `user.getWeeklyChartList`
- [x] `user.getWeeklyTrackChart`

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
