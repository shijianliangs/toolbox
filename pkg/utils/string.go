package utils

import "strings"

// StringUtils provides string manipulation utilities
type StringUtils struct{}

// NewStringUtils creates a new StringUtils instance
func NewStringUtils() *StringUtils {
	return &StringUtils{}
}

// Reverse reverses a string
func (s *StringUtils) Reverse(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToUpper converts string to uppercase
func (s *StringUtils) ToUpper(input string) string {
	return strings.ToUpper(input)
}

// ToLower converts string to lowercase
func (s *StringUtils) ToLower(input string) string {
	return strings.ToLower(input)
}

// TrimSpace trims whitespace from string
func (s *StringUtils) TrimSpace(input string) string {
	return strings.TrimSpace(input)
}

// Contains checks if string contains substring
func (s *StringUtils) Contains(input, substr string) bool {
	return strings.Contains(input, substr)
}

// Split splits string by delimiter
func (s *StringUtils) Split(input, delimiter string) []string {
	return strings.Split(input, delimiter)
}

// Join joins string slice with delimiter
func (s *StringUtils) Join(elements []string, delimiter string) string {
	return strings.Join(elements, delimiter)
}
