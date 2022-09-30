package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// uppercaseFile reads a file and returns its content in uppercase with included dependencies
func uppercaseFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file: %v", err)
	}

	return strings.ToUpper(string(content)), nil
}

// uppercase reads a reader and returns its content in uppercase with passed dependency
func uppercase(r io.Reader) (string, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read content: %v", err)
	}

	return strings.ToUpper(string(content)), nil
}
