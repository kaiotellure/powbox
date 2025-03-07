package tools

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var (
	LINK_FORMAT = "\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// Filters non-recursively files in the directory with the provided extension.
// Provide a dot prefixed extension (eg: .png | .txt)
func Files(dir, ext string) (list []string) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	// Iterate over directory entries
	for _, entry := range entries {
		// Skip directories, only process files
		if !entry.IsDir() {
			// Check if the file has a .txt extension
			if filepath.Ext(entry.Name()) == ext {
				// Append the full path or just the filename, depending on your need
				fullPath := filepath.Join(dir, entry.Name())
				list = append(list, fullPath)
			}
		}
	}

	return
}

// Creates a bufio reader from the file in the path.
// Remember to defer close the file!
// Panics on error!
func Reader(filepath string) (*bufio.Reader, *os.File) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	return bufio.NewReader(file), file
}

// Creates or opens a write only appendable file.
// Remember to defer close the file!
// Panics on error!
func Appendable(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	return file
}

// Returns a filename with the current date (eg: prefix-2006-01-02-15-04-05.ext)
// Provide a dot prefixed extension (eg: .png | .txt)
func TimedFilename(prefix, ext string) string {
	return fmt.Sprintf("%s-%s%s", prefix, time.Now().Format("2006-01-02-15-04-05"), ext)
}

var (
	spinnerSprites       = []string{".  ", ".. ", "..."}
	spinnerSpritesLength = len(spinnerSprites)
)

// Returns a string with the current spinner sprite for the current time, call repeatedly for a nice animation
func Spinner() string {
	// time slower and capped up to spinnerSpritesLength
	index := (int(time.Now().Unix()) % spinnerSpritesLength)
	return spinnerSprites[index]
}

// Reads input from user using a bufio scanner
func Ask() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

// Returns a list containing the keys of a map
func Keys[K string, V any](m map[K]V) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

// Ensures a path to a file or folder exists, creating directories that doesn't exists
func Ensure(path string) {
	os.MkdirAll(Dir(path), os.ModePerm)
}

// Returns the parent directory of a path. For a folder ending with a trailing path separator (/, \) such as (C://folder/) it returns itself
func Dir(path string) string {
	forward := filepath.ToSlash(path)
	return filepath.Dir(forward)
}