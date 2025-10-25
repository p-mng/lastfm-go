package lastfmgo_test

import (
	"testing"

	lastfmgo "github.com/p-mng/lastfm-go"
	"github.com/stretchr/testify/require"
)

func TestNewWebClient(t *testing.T) {
	_, err := lastfmgo.NewWebClient(lastfmgo.BaseURL, testHash, testHash, "http://localhost:8080")
	require.NoError(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, "", testHash, "http://localhost:8080")
	require.Error(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, testHash, "", "http://localhost:8080")
	require.Error(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, testHash, testHash, "%")
	require.Error(t, err)
}

func TestNewMobileClient(t *testing.T) {
	_, err := lastfmgo.NewMobileClient(lastfmgo.BaseURL, testHash, testHash, "username", "password")
	require.NoError(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, "", testHash, "username", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, testHash, "", "username", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, testHash, testHash, "", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, testHash, testHash, "username", "")
	require.Error(t, err)
}

func TestNewDesktopClient(t *testing.T) {
	_, err := lastfmgo.NewDesktopClient(lastfmgo.BaseURL, testHash, testHash)
	require.NoError(t, err)

	_, err = lastfmgo.NewDesktopClient(lastfmgo.BaseURL, "", testHash)
	require.Error(t, err)

	_, err = lastfmgo.NewDesktopClient(lastfmgo.BaseURL, testHash, "")
	require.Error(t, err)
}
