/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package utils

import (
	"net/url"
)

// URLEncode encodes a string to URL format
func URLEncode(input string) string {
	return url.QueryEscape(input)
}

// URLDecode decodes a URL format string to original string
func URLDecode(encoded string) (string, error) {
	return url.QueryUnescape(encoded)
}
