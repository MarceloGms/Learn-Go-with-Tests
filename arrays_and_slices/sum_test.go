package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	/* t.Run("Array", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		AssertCorrectMessage(t, got, expected, numbers)
	}) */
	t.Run("Slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		AssertCorrectMessage(t, got, expected, numbers)
	})
}

func AssertCorrectMessage(t testing.TB, got, expected int, numbers []int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d, but expected %d given %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{0,9})
	want := []int{3,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, but expected %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	//! new: assign a function to a variable
	checkSums := func (t testing.TB, got, want []int)  {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, but expected %v", got, want)
		}
	}
	t.Run("sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}
		
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0,9})
		want := []int{0,9}
	
		checkSums(t, got, want)
	})
}