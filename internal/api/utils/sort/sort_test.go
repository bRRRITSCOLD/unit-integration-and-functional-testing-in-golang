package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortAsc(t *testing.T) {
	// Init
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	// Execution
	BubbleSort(elements)

	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkBubbleSortAsc(b *testing.B) {
	// Init
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		// Execution
		BubbleSort(elements)
	}
}

func TestSortAsc(t *testing.T) {
	// Init
	elements := GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	// Execution
	Sort(elements)

	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func BenchmarkSortAsc(b *testing.B) {
	// Init
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		// Execution
		Sort(elements)
	}
}
