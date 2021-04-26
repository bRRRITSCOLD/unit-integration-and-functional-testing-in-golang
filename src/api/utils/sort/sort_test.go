package sort

import (
	"testing"
)

func TestBubbleSortAsc(t *testing.T) {
	// Init
	elements := GetElements(10000000)

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}

func BenchmarkBubbleSortAsc(b *testing.B) {
	// Init
	elements := GetElements(10000000)

	for i := 0; i < b.N; i++ {
		// Execution
		BubbleSort(elements)
	}
}

func TestSortAsc(t *testing.T) {
	// Init
	elements := GetElements(10000000)

	// Execution
	Sort(elements)

	// Validation
	if elements[0] != 0 {
		t.Error("first element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last element should be 9")
	}
}

func BenchmarkSortAsc(b *testing.B) {
	// Init
	elements := GetElements(10000000)

	for i := 0; i < b.N; i++ {
		// Execution
		Sort(elements)
	}
}
