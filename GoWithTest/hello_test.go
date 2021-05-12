package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	Want := "Hello, World"

	if got != Want {
		t.Errorf("Got %q want %q", got, Want)
	}
}
