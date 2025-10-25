package lastfmgo

import (
	"errors"
	"fmt"
	"strings"
)

// BaseURL is the default base URL of the last.fm API.
const BaseURL = "https://ws.audioscrobbler.com/2.0/"

// P is a shorthand for map[string]any.
type P map[string]any

// BaseResponse contains fields found in all responses.
type BaseResponse struct {
	Status string `xml:"status,attr"`
	Error  struct {
		Code    int64  `xml:"code,attr"`
		Message string `xml:",chardata"`
	} `xml:"error"`
}

// PaginationResponse contains fields found in paginated responses.
type PaginationResponse struct {
	Page       int64 `xml:"page,attr"`
	PerPage    int64 `xml:"perPage,attr"`
	Total      int64 `xml:"total,attr"`
	TotalPages int64 `xml:"totalPages,attr"`
}

// OpenSearchResponse contains fields found in OpenSearch query responses.
type OpenSearchResponse struct {
	For   string `xml:"for,attr"`
	Query struct {
		Role        string `xml:"role,attr"`
		SearchTerms string `xml:"searchTerms,attr"`
		StartPage   int64  `xml:"startPage,attr"`
	} `xml:"Query"`
	TotalResults int64 `xml:"totalResults"`
	StartIndex   int64 `xml:"startIndex"`
	ItemsPerPage int64 `xml:"itemsPerPage"`
}

type getError interface {
	getError() error
}

func (b BaseResponse) getError() error {
	switch b.Status {
	case "ok":
		return nil
	case "failed":
		return fmt.Errorf("%s (code %d)", strings.TrimSpace(b.Error.Message), b.Error.Code)
	default:
		return errors.New("API returned invalid status (must be ok or failed)")
	}
}
