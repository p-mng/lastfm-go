//nolint:revive
package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/michiwend/gomusicbrainz"
	lastfmgo "github.com/p-mng/lastfm-go"
)

func main() {
	_ = godotenv.Load()

	client, err := lastfmgo.NewMobileClient(
		lastfmgo.BaseURL,
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		os.Getenv("API_USER"),
		os.Getenv("API_PASS"),
	)
	if err != nil {
		panic(err.Error())
	}

	info, err := client.ArtistGetInfo(lastfmgo.P{
		"artist": "Red Hot Chili Peppers",
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Artist Name:", info.Artist.Name)
	fmt.Println("MBID:", info.Artist.MBID)

	mbClient, err := gomusicbrainz.NewWS2Client(
		"https://musicbrainz.org/ws/2",
		"lastfm-go",
		"1.0.0",
		"https://github.com/p-mng/lastfm-go",
	)
	if err != nil {
		panic(err.Error())
	}

	artist, err := mbClient.LookupArtist(gomusicbrainz.MBID(info.Artist.MBID))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Information from MusicBrainz:")
	fmt.Println(" - Formed:", artist.Lifespan.Begin.Format("2006"))
	fmt.Println(" - Location:", artist.Area.Name)
}
