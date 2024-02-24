package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Phil")
	want := "Hello, Phil"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
