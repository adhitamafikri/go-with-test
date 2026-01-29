package hello_world

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people in English", func(t *testing.T) {
		got := Hello("Fikri", "en")
		want := "Hello, Fikri"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' if empty string is supplied", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Sofie Reyes", "es")
		want := "Hola, Sofie Reyes"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello to people in French", func(t *testing.T) {
		got := Hello("Louis", "fr")
		want := "Alo, Louis"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
