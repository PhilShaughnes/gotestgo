package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to provided name", func(t *testing.T) {
		got := Hello("Phil", english)
		want := "Hello, Phil"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say 'Hello, World' when no name is provided", func(t *testing.T) {
		got := Hello("", english)
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("use french", func(t *testing.T) {
		got := Hello("", french)
		want := "Bonjour, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("use spanish", func(t *testing.T) {
		got := Hello("", spanish)
		want := "Hola, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
