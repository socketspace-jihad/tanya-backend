package middlewares

import (
	"errors"
	"strings"
)

func tokenParser(token string) (string, error) {
	if token == "" {
		return "", errors.New("authorization header cannot be nil")
	}
	parsed := strings.Split(token, " ")
	if len(parsed) < 2 {
		return "", errors.New("invalid token format")
	}
	return parsed[1], nil
}
