package lastfmgo

import (
	"encoding/xml"

	"io"
	"net/http"
	"strings"

	"github.com/p-mng/lastfm-go/internal"
)

// GetRequest sends an HTTP GET request to the last.fm API and decodes the response in the provided generic type.
//
// https://www.last.fm/api/rest
func GetRequest[T getError](
	baseURL string,
	params P,
	method string,
	apiKey string,
	apiSecret string,
	addSignature bool,
) (T, error) {
	var decoded T

	params["method"] = method
	params["api_key"] = apiKey

	if addSignature {
		signature, err := internal.CreateSignature(params, apiSecret)
		if err != nil {
			return decoded, err
		}
		params["api_sig"] = internal.EncodeSignature(signature)
	}

	url, err := internal.CreateURL(baseURL, params)
	if err != nil {
		return decoded, err
	}

	//nolint:gosec
	res, err := http.Get(url)
	if err != nil {
		return decoded, err
	}
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(res.Body)

	if err := xml.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return decoded, err
	}

	if err := decoded.getError(); err != nil {
		return decoded, err
	}

	return decoded, nil
}

// PostRequest sends an HTTP POST request to the last.fm API and decodes the response in the provided generic type.
//
// https://www.last.fm/api/rest
func PostRequest[T getError](
	baseURL string,
	params P,
	method string,
	apiKey string,
	apiSecret string,
	addSignature bool,
) (T, error) {
	var decoded T

	params["method"] = method
	params["api_key"] = apiKey

	if addSignature {
		signature, err := internal.CreateSignature(params, apiSecret)
		if err != nil {
			return decoded, err
		}
		params["api_sig"] = internal.EncodeSignature(signature)
	}

	encodedParams, err := internal.EncodeParams(params)
	if err != nil {
		return decoded, err
	}

	req, err := http.NewRequest(http.MethodPost, baseURL, strings.NewReader(encodedParams))
	if err != nil {
		return decoded, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return decoded, err
	}
	defer func(r io.ReadCloser) {
		_ = r.Close()
	}(res.Body)

	if err := xml.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return decoded, err
	}

	if err := decoded.getError(); err != nil {
		return decoded, err
	}

	return decoded, nil
}
