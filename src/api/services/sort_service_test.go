package services

import "testing"

func TestSort(t *testing.T) {
	// Init
	elements := []int{9, 3, 4, 1, 6, 7, 5, 2, 8, 0}

	Sort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}
