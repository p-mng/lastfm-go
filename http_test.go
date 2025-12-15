package lastfmgo_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	lastfmgo "github.com/p-mng/lastfm-go"
	"github.com/stretchr/testify/require"
)

//nolint:gosec
const (
	TestKey    = "6e87a4f6ea8244d7e04cb52e495d6693"
	TestSecret = "d9729feb74992cc3482b350163a1a010"
)

func TestGetRequest(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		data, err := os.ReadFile("testdata/geo.getTopArtists.xml")
		require.NoError(t, err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			require.Equal(t, "geo.getTopArtists", r.URL.Query().Get("method"))
			require.Equal(t, TestKey, r.URL.Query().Get("api_key"))
			require.Equal(t, "", r.URL.Query().Get("api_sig"))
			require.Equal(t, "", r.URL.Query().Get("sk"))

			require.Equal(t, "Spain", r.URL.Query().Get("country"))
			require.Equal(t, "50", r.URL.Query().Get("limit"))
			require.Equal(t, "1", r.URL.Query().Get("page"))

			w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}))
		defer server.Close()

		response, err := lastfmgo.GetRequest[lastfmgo.GeoGetTopArtistsResponse](
			server.URL,
			lastfmgo.P{
				"country": "Spain",
				"limit":   50,
				"page":    int64(1),
			},
			"geo.getTopArtists",
			TestKey,
			TestSecret,
			false,
		)

		require.NoError(t, err)
		require.Equal(t, "ok", response.Status)
	})

	t.Run("Invalid Request", func(t *testing.T) {
		data, err := os.ReadFile("testdata/geo.getTopArtists_error.xml")
		require.NoError(t, err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodGet, r.Method)
			require.Equal(t, "geo.getTopArtists", r.URL.Query().Get("method"))

			w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}))
		defer server.Close()

		response, err := lastfmgo.GetRequest[lastfmgo.GeoGetTopArtistsResponse](
			server.URL,
			lastfmgo.P{
				"country": "Nambutu",
				"limit":   50,
				"page":    int64(1),
			},
			"geo.getTopArtists",
			TestKey,
			TestSecret,
			false,
		)

		require.Error(t, err)
		require.Equal(t, "failed", response.Status)
	})

	t.Run("HTTP Error", func(t *testing.T) {
		_, err := lastfmgo.GetRequest[lastfmgo.GeoGetTopArtistsResponse](
			"http://invalid:60000",
			lastfmgo.P{
				"country": "Spain",
				"limit":   50,
				"page":    int64(1),
			},
			"geo.getTopArtists",
			TestKey,
			TestSecret,
			false,
		)
		require.Error(t, err)
	})
}

func TestPostRequest(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		data, err := os.ReadFile("testdata/track.love.xml")
		require.NoError(t, err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodPost, r.Method)
			require.Len(t, r.URL.Query(), 0)

			body, err := io.ReadAll(r.Body)
			require.NoError(t, err)

			parts := strings.Split(string(body), "&")
			require.Equal(t, parts, []string{
				"api_key=6e87a4f6ea8244d7e04cb52e495d6693",
				"api_sig=55e6b6cc9f5c2c3838bd924b933e7c61",
				"artist=Radiohead",
				"method=track.love",
				"track=Fake+Plastic+Trees",
			})

			w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}))
		defer server.Close()

		response, err := lastfmgo.PostRequest[lastfmgo.TrackLoveResponse](
			server.URL,
			lastfmgo.P{
				"track":  "Fake Plastic Trees",
				"artist": "Radiohead",
			},
			"track.love",
			TestKey,
			TestSecret,
			true,
		)

		require.NoError(t, err)
		require.Equal(t, "ok", response.Status)
	})

	t.Run("Invalid Request", func(t *testing.T) {
		data, err := os.ReadFile("testdata/track.love_error.xml")
		require.NoError(t, err)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			require.Equal(t, http.MethodPost, r.Method)
			require.Len(t, r.URL.Query(), 0)

			w.Header().Set("Content-Type", "text/xml; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
		}))
		defer server.Close()

		response, err := lastfmgo.PostRequest[lastfmgo.TrackLoveResponse](
			server.URL,
			lastfmgo.P{
				"track":  "Fake Plastic Trees",
				"artist": "Radiohead",
			},
			"track.love",
			TestKey,
			TestSecret,
			true,
		)

		require.Error(t, err)
		require.Equal(t, "failed", response.Status)
	})

	t.Run("HTTP Error", func(t *testing.T) {
		_, err := lastfmgo.PostRequest[lastfmgo.TrackLoveResponse](
			"http://invalid:60000",
			lastfmgo.P{
				"track":  "Fake Plastic Trees",
				"artist": "Radiohead",
			},
			"track.love",
			TestKey,
			TestSecret,
			true,
		)

		require.Error(t, err)
	})
}
