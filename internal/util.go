// Package internal provides basic functions for building last.fm API requests.
package internal

import (
	//nolint:gosec
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"slices"
	"strconv"
)

// CreateSignature creates an unencoded signature using the provided parameters and the API secret.
// Requires calling [EncodeSignature] before attaching to API calls.
func CreateSignature(params map[string]any, secret string) (string, error) {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	var signature string
	for _, key := range keys {
		signature += key

		value, err := ToString(params[key])
		if err != nil {
			return "", err
		}
		signature += value
	}
	signature += secret

	return signature, nil
}

// EncodeSignature encodes the provided signature using MD5.
func EncodeSignature(signature string) string {
	//nolint:gosec
	hash := md5.Sum([]byte(signature))
	return hex.EncodeToString(hash[:])
}

// CreateURL encodes the provided parameters and creates a URL for HTTP GET requests.
func CreateURL(baseURL string, params map[string]any) (string, error) {
	paramsEncoded, err := EncodeParams(params)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s?%s", baseURL, paramsEncoded), nil
}

// EncodeParams URL-encodes the provided parameters for usage with [CreateURL] or as a body in HTTP POST requests.
func EncodeParams(params map[string]any) (string, error) {
	query := url.Values{}
	for k, v := range params {
		value, err := ToString(v)
		if err != nil {
			return "", err
		}
		query.Add(k, value)
	}
	return query.Encode(), nil
}

// ToString converts the given value into a string.
//
// Returns an error if the type of value is not supported.
func ToString(input any) (string, error) {
	switch input := input.(type) {
	case string:
		return input, nil
	case int:
		return strconv.Itoa(input), nil
	case int64:
		return strconv.FormatInt(input, 10), nil
	default:
		return "", errors.New("failed to convert an unknown type to string")
	}
}

// ToInt converts the given boolean into an integer.
func ToInt(input bool) int {
	if input {
		return 1
	}
	return 0
}

// IsMD5 checks if the provided string is a valid MD5 hash.
func IsMD5(input string) bool {
	re := regexp.MustCompile("^[a-f0-9]{32}$")
	return re.MatchString(input)
}

// IsURL checks if the provided string is a valid URL.
func IsURL(input string) bool {
	_, err := url.Parse(input)
	return err == nil
}
