//nolint:revive
package main

import (
	"fmt"

	lastfmgo "github.com/p-mng/lastfm-go"
)

func main() {
	client, err := lastfmgo.NewMobileClient(
		lastfmgo.BaseURL,
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"InvalidUsername",
		"InvalidPassword",
	)
	if err != nil {
		panic(err.Error())
	}

	_, err = client.AuthGetMobileSession()
	fmt.Println("Error:", err.Error())
}
