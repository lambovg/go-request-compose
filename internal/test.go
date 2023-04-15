package test

import (
	"reflect"
	"regexp"
	"testing"
)

func Ok(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Equals(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func Match(t *testing.T, got string, want string) {
	m, _ := regexp.MatchString(want, got)

	if !m {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
