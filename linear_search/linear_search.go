package linear_search

//Sequential Search && Linear Search
//循序 & 線性搜尋演算法
func LinearSearch(array []int, foundValue int) int {
	for index, value := range array {
		if foundValue == value {
			return index
		}
	}

	return -1
}
