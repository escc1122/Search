package binary_search

//二分搜尋演算法
func BinarySearch(array []int, left int, right int, foundValue int) int {
	if left > right {
		return -1
	}

	middle := (int)(left+right) / 2
	middleValue := array[middle]

	if middleValue == foundValue {
		return middle
	} else if foundValue > middleValue {
		return BinarySearch(array, middle+1, right, foundValue)
	} else if foundValue < middleValue {
		return BinarySearch(array, left, middle-1, foundValue)
	}
	return -1
}

//插值搜尋演算法
func InterpolationSearch(array []int, left int, right int, foundValue int) int {
	if left > right {
		return -1
	}
	if foundValue > array[right] {
		return -1
	}

	if foundValue < array[left] {
		return -1
	}

	// middle := (int)(left+right) / 2
	//插值運算中心點公式
	middle := left + (right-left)*(foundValue-array[left])/(array[right]-array[left])
	middleValue := array[middle]

	if middleValue == foundValue {
		return middle
	} else if foundValue > middleValue {
		return InterpolationSearch(array, middle+1, right, foundValue)
	} else if foundValue < middleValue {
		return InterpolationSearch(array, left, middle-1, foundValue)
	}
	return -1
}
