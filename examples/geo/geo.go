//nolint:revive
package main

import (
	"fmt"
	"os"

	"github.com/biter777/countries"
	"github.com/joho/godotenv"
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

	topArtists, err := client.GeoGetTopArtists(lastfmgo.P{
		"country": countries.Spain.String(),
		"limit":   5,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Top artists in Spain:")
	for _, artist := range topArtists.TopArtists.Artists {
		fmt.Printf(" - %s (listeners: %d)\n", artist.Name, artist.Listeners)
	}

	topTracks, err := client.GeoGetTopTracks(lastfmgo.P{
		"country": countries.UnitedStatesOfAmerica.String(),
		"limit":   5,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Top tracks in the United States:")
	for _, track := range topTracks.Tracks.Tracks {
		fmt.Printf(" - %s - %s (listeners: %d)\n", track.Artist.Name, track.Name, track.Listeners)
	}
}
