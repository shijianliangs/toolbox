package utils

import (
	"os"
	"path/filepath"
)

// FileUtils provides file operation utilities
type FileUtils struct{}

// NewFileUtils creates a new FileUtils instance
func NewFileUtils() *FileUtils {
	return &FileUtils{}
}

// Exists checks if file or directory exists
func (f *FileUtils) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// IsDir checks if path is a directory
func (f *FileUtils) IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// Size returns the size of a file or directory
func (f *FileUtils) Size(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		var size int64
		err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				size += info.Size()
			}
			return nil
		})
		return size, err
	}

	return info.Size(), nil
}

// CreateDir creates a directory if it doesn't exist
func (f *FileUtils) CreateDir(path string) error {
	if !f.Exists(path) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}
