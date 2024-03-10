package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth reader")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid auth type: " + vals[0])
	}

	return vals[1], nil
}
