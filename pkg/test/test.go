package test

import "testing"

func Ok(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
