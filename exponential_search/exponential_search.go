package exponential_search

import (
	"search/binary_search"
)

//指數搜尋演算法
func ExponentialSearch(array []int, foundValue int) int {
	bound := 1

	for bound < len(array) && foundValue > array[bound] {
		bound = bound * 2
	}
	var right int
	if bound > len(array) {
		right = len(array)
	} else {
		right = bound
	}
	index := binary_search.BinarySearch(array, bound/2, right-1, foundValue)
	return index
}
