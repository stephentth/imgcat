package main

import (
	"net/url"
)

func IsValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}
	return true
}

func IsValidFile(filename string) bool {
	return true
}
