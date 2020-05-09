package utilities

import (
	"testing"
)

func TestIntSliceIsEqualWithEqualEmptySlices(t *testing.T) {
	given := IntSlice([]int{})
	expected := true
	actual := given.IsEqualTo([]int{})
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithUnequalEmptySlices(t *testing.T) {
	given := IntSlice([]int{})
	expected := false
	actual := given.IsEqualTo([]int{1})
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithEqualSlices(t *testing.T) {
	given := IntSlice([]int{1, 2, 3})
	expected := true
	actual := given.IsEqualTo([]int{1, 2, 3})
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithUnequalSlices(t *testing.T) {
	given := IntSlice([]int{1, 2, 3})
	expected := false
	actual := given.IsEqualTo([]int{3, 2, 1})
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}

func TestIntSliceIsEqualWithDifferentLengthSlices(t *testing.T) {
	given := IntSlice([]int{1, 2, 3})
	expected := false
	actual := given.IsEqualTo([]int{1, 2})
	if expected != actual {
		t.Errorf("failed to get expected result: expected %v: got %v", expected, actual)
	}
}
