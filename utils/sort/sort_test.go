package sort

import (
	"testing"
)

func TestBubbleSortDesc(t *testing.T) {
	// Init
	elements := []int{9, 3, 4, 1, 6, 7, 5, 2, 8, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 1")
	}
}
