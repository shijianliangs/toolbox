package utils

import (
	"encoding/base64"
)

// Base64Utils provides base64 encoding and decoding utilities
type Base64Utils struct{}

// NewBase64Utils creates a new Base64Utils instance
func NewBase64Utils() *Base64Utils {
	return &Base64Utils{}
}

// Encode encodes data to base64 string
func (b *Base64Utils) Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode decodes base64 string to data
func (b *Base64Utils) Decode(encoded string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encoded)
}

// EncodeURL encodes data to URL-safe base64 string
func (b *Base64Utils) EncodeURL(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

// DecodeURL decodes URL-safe base64 string to data
func (b *Base64Utils) DecodeURL(encoded string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(encoded)
}
