package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, given %d", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 5})
	want := []int{3, 9}

	//useful for seeing if any two variables are the same, convenient way of comparing slices
	//important to note that reflect.DeepEqual is not "type safe"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 3})
		want := []int{2, 3}
		checkSums(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2})
		want := []int{0, 2}
		checkSums(t, got, want)
	})
}
