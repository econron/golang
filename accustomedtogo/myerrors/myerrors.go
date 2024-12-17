package myerrors

import (
	"net/http"
	"io"
)

type HTTPError struct {
	StatusCode int
	URL string
}

// これでerrorインターフェースを満たす
func (he HTTPError) Error() string {
	return "HTTP Error: " + he.URL
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, HTTPError{StatusCode: 500, URL: url}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, HTTPError{StatusCode: resp.StatusCode, URL: url}
	}

	return io.ReadAll(resp.Body)
}