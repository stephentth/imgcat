package main

import (
	"io"
	"net/http"
)

func GetImageFromHTTP(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}
