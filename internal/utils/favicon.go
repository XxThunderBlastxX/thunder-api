package utils

import (
	"io"
	"net/http"
)

func GetFavicon() ([]byte, error) {
	res, err := http.Get("https://koustav.dev/favicon.ico")
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
