package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to someone", func(t *testing.T) {
		got := Hello("Marcelo", "")
		want := "Hello, Marcelo"

		AssertCorrectMessage(t, got, want)
	})
	t.Run("Say 'Hello, world!' if empty string is provided", func (t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		AssertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Marcelo", "Spanish")
		want := "Hola, Marcelo"

		AssertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Marcelo", "French")
		want := "Bonjour, Marcelo"

		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}