package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"testing/fstest"
)

func Test_readWordsFromFile(t *testing.T) {
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("create test file: %v", err)
	}

	if _, err := fmt.Fprint(tf, "one\ntwo\tthree four      five"); err != nil {
		t.Fatalf("write test content to temp file: %v", err)
	}

	if err := tf.Close(); err != nil {
		t.Fatalf("close test file: %v", err)
	}

	words, err := readWordsFromFile(tf.Name())
	if err != nil {
		t.Fatalf("read test file: %v", err)
	}

	expected := []string{"one", "two", "three", "four", "five"}
	if !reflect.DeepEqual(words, expected) {
		t.Fatalf("result %+v != expected %+v", words, expected)
	}
}

func Test_readWordsFromFileV2(t *testing.T) {
	const name = "some/path"
	tfs := fstest.MapFS{name: &fstest.MapFile{
		Data: []byte("one\ntwo\tthree four      five"),
	}}

	words, err := readWordsFromFileV2(tfs, name)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"one", "two", "three", "four", "five"}
	if !reflect.DeepEqual(words, expected) {
		t.Fatalf("result %+v != expected %+v", words, expected)
	}
}
