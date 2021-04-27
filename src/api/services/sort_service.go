package services

import "unit-integration-and-functional-testing-in-golang/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) > 10000 {
		sort.Sort(elements)
		return
	}
	sort.BubbleSort(elements)
}
