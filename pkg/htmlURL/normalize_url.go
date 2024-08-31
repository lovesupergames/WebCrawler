package htmlURL

import (
	"errors"
	"net/url"
)

func NormalizeURL(RawUrl string) (string, error) {

	normalizedURL, err := url.Parse(RawUrl)
	if err != nil {
		return "", err
	}
	if len(normalizedURL.Host) == 0 {
		return "", errors.New("invalid URL")
	}
	if normalizedURL.Scheme != "http" && normalizedURL.Scheme != "https" {
		return "", errors.New("URL scheme must be http or https")
	}
	if normalizedURL.Path != "" && normalizedURL.Path[len(normalizedURL.Path)-1] == '/' {
		normalizedURL.Path = normalizedURL.Path[:len(normalizedURL.Path)-1]
	}

	return normalizedURL.Host + normalizedURL.Path, nil
}
