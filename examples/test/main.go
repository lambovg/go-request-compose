package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("Usage: words-cli ")
		return
	}

	fsProvider := os.DirFS("/")

	target := flag.Arg(0)
	words, err := readWordsFromFileV2(fsProvider, target)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, word := range words {
		fmt.Println(word)
	}
}

func readWordsFromFile(path string) ([]string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	return strings.Fields(string(content)), nil
}

func readWordsFromFileV2(f fs.FS, path string) ([]string, error) {
	content, err := fs.ReadFile(f, path)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	return strings.Fields(string(content)), nil
}
