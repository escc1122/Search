package main

import (
	"fmt"
	"math/rand"
	"search/binary_search"
	"search/exponential_search"
	"search/linear_search"
	"time"
)

func testLinearSearch() {
	rand.Seed(time.Now().Unix())
	test := rand.Perm(10)
	fmt.Println(test)
	index := linear_search.LinearSearch(test, 5)
	fmt.Println(index)
}

func testBinarySearch() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(test)
	index := binary_search.BinarySearch(test, 0, len(test)-1, 10)
	fmt.Println(index)
}

func testInterpolationSearch() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(test)
	fmt.Println(binary_search.InterpolationSearch(test, 0, len(test)-1, 0))
	fmt.Println(binary_search.InterpolationSearch(test, 0, len(test)-1, 5))
	fmt.Println(binary_search.InterpolationSearch(test, 0, len(test)-1, 9))
	fmt.Println(binary_search.InterpolationSearch(test, 0, len(test)-1, -1))
	fmt.Println(binary_search.InterpolationSearch(test, 0, len(test)-1, 11))

}

func testExponentialSearch() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(exponential_search.ExponentialSearch(test, 0))
	fmt.Println(exponential_search.ExponentialSearch(test, 5))
	fmt.Println(exponential_search.ExponentialSearch(test, 9))
	fmt.Println(exponential_search.ExponentialSearch(test, -1))
	fmt.Println(exponential_search.ExponentialSearch(test, 11))
}

func main() {
	// testLinearSearch()
	// testBinarySearch()
	testInterpolationSearch()
	// testExponentialSearch()
}
