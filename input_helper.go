package main

import (
	"net/url"
)

// IsValidURL check if a input is a URL
func IsValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	return true
}
