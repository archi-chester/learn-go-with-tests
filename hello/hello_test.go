package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Chester", "Spanish")
		want := "Hola, Chester"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Chester", "French")
		want := "Bonjour, Chester"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chester", "")
		want := "Hello, Chester"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' if an empty string was sent", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}
