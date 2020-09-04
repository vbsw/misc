/*
 *          Copyright 2020, Vitali Baumtrok.
 * Distributed under the Boost Software License, Version 1.0.
 *     (See accompanying file LICENSE or copy at
 *        http://www.boost.org/LICENSE_1_0.txt)
 */

// Package files provides functions to check file status.
package files

import (
	"os"
	"path/filepath"
)

// Exists returns true, if file exists. It doesn't matter if it is a directory or a file.
func Exists(path string) bool {
	fileInfo, err := os.Stat(path)
	return fileInfo != nil && (err == nil || !os.IsNotExist(err))
}

// IsDirectory returns true, if file exists and is a directory.
func IsDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	return fileInfo != nil && fileInfo.IsDir() && (err == nil || !os.IsNotExist(err))
}

// IsFile returns true, if file exists and is not a directory.
func IsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	return fileInfo != nil && !fileInfo.IsDir() && (err == nil || !os.IsNotExist(err))
}

// Size returns size of file in bytes. It doesn't matter if it is a directory or a file.
// If file is a directory all subdirectories are included.
// If file does not exist it returns 0.
func Size(path string) int64 {
	fileInfo, err := os.Stat(path)
	if fileInfo != nil && (err == nil || !os.IsNotExist(err)) {
		if !fileInfo.IsDir() {
			return FileSize(path)
		}
		return DirSizeR(path)
	}
	return 0
}

// FileSize returns size of file in bytes. If file does not exist or is a directory it returns 0.
func FileSize(path string) int64 {
	fileInfo, err := os.Stat(path)
	if fileInfo != nil && !fileInfo.IsDir() && (err == nil || !os.IsNotExist(err)) {
		return fileInfo.Size()
	}
	return 0
}

// DirSizeR returns total size of files in bytes, including subdirectories.
// If directory does not exist or is a file it returns 0.
func DirSizeR(path string) int64 {
	var size int64
	filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err == nil {
			size += FileSize(filePath)
		}
		return err
	})
	return size
}

// DirSizeF returns total size of files in bytes, no subdirectories included.
// If directory does not exist or is a file it returns 0.
func DirSizeF(path string) int64 {
	var size int64
	filepath.Walk(path, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err == nil {
			size += FileSize(filePath)
		}
		return err
	})
	return size
}
