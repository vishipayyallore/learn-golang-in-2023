package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {

		// t.Helper() is needed to tell the test suite that this method is a helper.
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sri Varu", "English")
		want := "Hello, Sri Varu"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// Additional Notes:
// t.Fail() is a function that can be called to indicate that a test has failed.

// got := Hello("Sri Varu")
// want := "Hello, Sri Varu"

// if got != want {
// 	t.Errorf("got %q want %q", got, want)
// }
