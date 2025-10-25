package internal_test

import (
	"testing"

	"github.com/p-mng/lastfm-go/internal"
	"github.com/stretchr/testify/require"
)

func TestCreateSignature(t *testing.T) {
	sigExpected := "api_key91856cb7960b9f7751fd4a16c7bb3492methodauth.getSessionsomeInteger597a7bba8b9be470c7f4b8760fa4086af0"
	sigActual, err := internal.CreateSignature(map[string]any{
		"method":      "auth.getSession",
		"api_key":     "91856cb7960b9f7751fd4a16c7bb3492",
		"someInteger": 59,
	}, "7a7bba8b9be470c7f4b8760fa4086af0")

	require.NoError(t, err)
	require.Equal(t, sigExpected, sigActual)
}

func TestEncodeSignature(t *testing.T) {
	sig := "api_keyGEYotGy8yWCqrYNgeSxabNQsZ2ztDzRwmethodauth.getSession7a7bba8b9be470c7f4b8760fa4086af0"

	require.Equal(t, "22b7af29e98e2e284206a64b1868d6fb", internal.EncodeSignature(sig))
}

func TestCreateURL(t *testing.T) {
	params := map[string]any{
		"method":   "user.getRecentTracks",
		"limit":    190,
		"user":     "JohnDoe",
		"page":     2,
		"extended": 1,
		"api_key":  "GEYotGy8yWCqrYNgeSxabNQsZ2ztDzRw",
	}

	urlExpected := "https://ws.audioscrobbler.com/2.0/?api_key=GEYotGy8yWCqrYNgeSxabNQsZ2ztDzRw&extended=1&limit=190&method=user.getRecentTracks&page=2&user=JohnDoe"
	urlActual, err := internal.CreateURL("https://ws.audioscrobbler.com/2.0/", params)

	require.NoError(t, err)
	require.Equal(t, urlExpected, urlActual)
}

func TestToString(t *testing.T) {
	// string
	valString, err := internal.ToString("hello, world")

	require.NoError(t, err)
	require.Equal(t, "hello, world", valString)

	// integer
	valInt, err := internal.ToString(43894)

	require.NoError(t, err)
	require.Equal(t, "43894", valInt)

	// unsupported type
	valError, err := internal.ToString(map[int]string{})

	require.Error(t, err)
	require.Equal(t, "", valError)
}

func TestToInt(t *testing.T) {
	require.Equal(t, 1, internal.ToInt(true))
	require.Equal(t, 0, internal.ToInt(false))
}

func TestIsMD5(t *testing.T) {
	require.True(t, internal.IsMD5("91856cb7960b9f7751fd4a16c7bb3492"))
	require.False(t, internal.IsMD5("lastfm-go"))
}

func TestIsValidURL(t *testing.T) {
	require.True(t, internal.IsURL("https://ws.audioscrobbler.com/2.0/"))
	require.False(t, internal.IsURL("%"))
}
