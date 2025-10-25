//nolint:revive
package main

import (
	"fmt"
	"os"

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

	fmt.Println("Authorized user:", session.Session.Name)
	fmt.Println("Session key:", session.Session.Key)
}
