package linear_search

import (
	"fmt"
	"testing"
)

func TestLinearSearch2(t *testing.T) {
	test := []int{0, 4, 7, 5, 6, 8, 2, 1, 3, 9}
	fmt.Println(test)
	index := LinearSearch(test, 5)
	fmt.Println(index)
	if index != 3 {
		t.Error()
	}
}
