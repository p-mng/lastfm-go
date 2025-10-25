//nolint:revive
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	lastfmgo "github.com/p-mng/lastfm-go"
)

func main() {
	_ = godotenv.Load()

	client, err := lastfmgo.NewDesktopClient(
		lastfmgo.BaseURL,
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
	)
	if err != nil {
		panic(err.Error())
	}

	token, err := client.AuthGetToken()
	if err != nil {
		panic(err.Error())
	}

	url := client.DesktopAuthorizationURL(token.Token)

	fmt.Println("Authorize app in browser:", url)
	fmt.Print("Press Enter to continue...")
	if _, err := bufio.NewReader(os.Stdin).ReadBytes('\n'); err != nil {
		panic(err.Error())
	}

	session, err := client.AuthGetSession(token.Token)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Authorized user:", session.Session.Name)
	fmt.Println("Session key:", session.Session.Key)
}
