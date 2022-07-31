package binary_search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(test)
	index := BinarySearch(test, 0, len(test)-1, 10)
	fmt.Println(index)
	if index != -1 {
		t.Error()
	}

}
