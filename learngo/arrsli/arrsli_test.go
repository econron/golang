package arrsli

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}

	expected := Sum(numbers)
	actual := 15

	if expected != actual {
		t.Errorf("expected '%d', actual '%d'", expected, actual)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,2}, []int{3,4})
	want := []int{3,7}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1,2}, []int{0,9})
	want := []int{2,9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	got2 := SumAllTails(nil, []int{0,9})
	want2 := []int{0,9}

	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %v want %v", got2, want2)
	}

}