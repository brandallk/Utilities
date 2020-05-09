package utilities

import (
	"testing"
)

func TestIntSliceIsEqualWithEqualEmptySlices(t *testing.T) {
	given1 := IntSlice([]int{})
	given2 := IntSlice([]int{})
	expected := true
	actual := given1.IsEqual(given2)
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithUnequalEmptySlices(t *testing.T) {
	given1 := IntSlice([]int{})
	given2 := IntSlice([]int{1})
	expected := false
	actual := given1.IsEqual(given2)
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithEqualSlices(t *testing.T) {
	given1 := IntSlice([]int{1, 2, 3})
	given2 := IntSlice([]int{1, 2, 3})
	expected := true
	actual := given1.IsEqual(given2)
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithUnequalSlices(t *testing.T) {
	given1 := IntSlice([]int{1, 2, 3})
	given2 := IntSlice([]int{3, 2, 1})
	expected := false
	actual := given1.IsEqual(given2)
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithDifferentLengthSlices(t *testing.T) {
	given1 := IntSlice([]int{1, 2, 3})
	given2 := IntSlice([]int{1, 2})
	expected := false
	actual := given1.IsEqual(given2)
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}
