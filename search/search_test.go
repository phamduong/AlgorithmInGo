package search

import "testing"

func TestLinearSearch(t *testing.T) {
	s := []int{1, 3, 4, 2}
	var find int = LinearSearch(s, 4, 2)
	if find != 3 {
		t.Errorf("Linear search error")
	}
}

func TestBinarySearch(t *testing.T) {
	s := []int{1, 3, 5, 9, 45, 234}
	var find = BinarySearch(s, 6, 5)
	if find != 2 {
		t.Error("Binary search error")
	}
}
