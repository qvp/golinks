package parser

import (
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

// LoadHtml fetches the HTML content from the specified URL
func LoadHtml(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error().Msgf("LoadHtml request fail: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Msgf("LoadHtml read body fail: %s\n", err)
		return "", err
	}

	return string(body), nil
}
