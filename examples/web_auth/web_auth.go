//nolint:revive
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	lastfmgo "github.com/p-mng/lastfm-go"
)

func main() {
	_ = godotenv.Load()

	client, err := lastfmgo.NewWebClient(
		lastfmgo.BaseURL,
		os.Getenv("API_KEY"),
		os.Getenv("API_SECRET"),
		"http://localhost:8080",
	)
	if err != nil {
		panic(err.Error())
	}

	url := client.WebAuthorizationURL()

	fmt.Println("Authorize app in browser:", url)

	tokenChan := make(chan string)

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", func(_ http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		tokenChan <- token
	})

	go func() {
		//nolint:gosec
		_ = http.ListenAndServe(":8080", serveMux)
	}()

	token := <-tokenChan

	session, err := client.AuthGetSession(token)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Logged in user name:", session.Session.Name)
	fmt.Println("Session key:", session.Session.Key)
}
