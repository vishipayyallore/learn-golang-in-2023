package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sri Varu")
		want := "Hello, Sri Varu"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	// got := Hello("Sri Varu")
	// want := "Hello, Sri Varu"

	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// }
}

// Additional Notes:
// t.Fail() is a function that can be called to indicate that a test has failed.
