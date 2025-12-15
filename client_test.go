package lastfmgo_test

import (
	"testing"

	lastfmgo "github.com/p-mng/lastfm-go"
	"github.com/stretchr/testify/require"
)

func TestNewWebClient(t *testing.T) {
	_, err := lastfmgo.NewWebClient(lastfmgo.BaseURL, TestKey, TestSecret, "http://localhost:8080")
	require.NoError(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, "", TestSecret, "http://localhost:8080")
	require.Error(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, TestKey, "", "http://localhost:8080")
	require.Error(t, err)

	_, err = lastfmgo.NewWebClient(lastfmgo.BaseURL, TestKey, TestSecret, "%")
	require.Error(t, err)
}

func TestNewMobileClient(t *testing.T) {
	_, err := lastfmgo.NewMobileClient(lastfmgo.BaseURL, TestKey, TestSecret, "username", "password")
	require.NoError(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, "", TestSecret, "username", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, TestKey, "", "username", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, TestKey, TestSecret, "", "password")
	require.Error(t, err)

	_, err = lastfmgo.NewMobileClient(lastfmgo.BaseURL, TestKey, TestSecret, "username", "")
	require.Error(t, err)
}

func TestNewDesktopClient(t *testing.T) {
	_, err := lastfmgo.NewDesktopClient(lastfmgo.BaseURL, TestKey, TestSecret)
	require.NoError(t, err)

	_, err = lastfmgo.NewDesktopClient(lastfmgo.BaseURL, "", TestSecret)
	require.Error(t, err)

	_, err = lastfmgo.NewDesktopClient(lastfmgo.BaseURL, TestKey, "")
	require.Error(t, err)
}
