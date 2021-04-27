package services

import (
	"testing"
	"unit-integration-and-functional-testing-in-golang/src/api/utils/sort"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	// Init
	elements := sort.GetElements(10)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 9, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	Sort(elements)

	assert.EqualValues(t, 10, len(elements))
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "last element should be 9")
}

func TestSortMoreThan10000(t *testing.T) {
	// Init
	elements := sort.GetElements(20001)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 20001, len(elements))
	assert.EqualValues(t, 20000, elements[0])
	assert.EqualValues(t, 0, elements[len(elements)-1])

	Sort(elements)

	assert.EqualValues(t, 20001, len(elements))
	assert.EqualValues(t, 0, elements[0], "first element should be 0")
	assert.EqualValues(t, 20000, elements[len(elements)-1], "last element should be 9")
}
