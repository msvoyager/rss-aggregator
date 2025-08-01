package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract a api key from heades from a HTTP request
// header format :-
// Authorization: ApiKey {insert apikey here}       -- Authorization: ApiKey abc123xyz
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")

	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}

	return vals[1], nil
}
