package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	res := Repeat("a", 5)
	fmt.Println(res)
	// Output: aaaaa
}
