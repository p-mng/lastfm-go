//nolint:revive
package main

import (
	"fmt"
	"os"
	"time"

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

	session, err := client.AuthGetMobileSession()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Sending scrobble to user:", session.Session.Name)

	timestamp := time.Now().Add(-1 * time.Minute).Unix()
	duration := 5*60 + 21

	// Find required and optional parameters here: https://www.last.fm/api/show/track.response
	response, err := client.TrackScrobble(lastfmgo.P{
		"artist":    "Red Hot Chili Peppers",
		"track":     "Californication",
		"timestamp": timestamp,
		"album":     "Californication",
		"duration":  duration,
		"sk":        session.Session.Key,
	})
	if err != nil {
		panic(err.Error())
	}

	scrobble := response.Scrobbles.Scrobbles[0]
	fmt.Println(" - Artist:", scrobble.Artist.Name)
	fmt.Println(" - Album:", scrobble.Album.Name)
	fmt.Println(" - Timestamp:", time.Unix(int64(scrobble.Timestamp), 0).Format(time.RFC1123))
}
