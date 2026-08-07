package security

import (
	"net/url"
	"strings"
)

func ValidateURL(value string) bool {

	value = strings.TrimSpace(value)

	if value == "" {
		return false
	}

	parsed, err := url.Parse(value)

	if err != nil {
		return false
	}

	return parsed.Scheme == "http" ||
		parsed.Scheme == "https"
}
