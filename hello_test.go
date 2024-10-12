package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say Hello to someone", func(t *testing.T) {
		got := Hello("Marcelo")
		want := "Hello, Marcelo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say 'Hello, world!' if empty string is provided", func (t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	
}