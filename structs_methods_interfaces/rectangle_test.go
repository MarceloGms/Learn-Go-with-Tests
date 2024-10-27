package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 10.0)
	want := 40.0

	MessageIf(t, got, want)
}

func TestArea(t *testing.T) {
	got := Area(10.0, 10.0)
	want := 100.0

	MessageIf(t, got, want)
}

func MessageIf(tb testing.TB, got, want float64) {
	tb.Helper()
	if got != want {
		tb.Errorf("got %.2f, want %.2f", got , want)
	}
}