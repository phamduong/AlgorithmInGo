package search

func LinearSearch(a []int, n int, x int) int {
	var i int = 0
	for i < n && x != a[i] {
		i++
	}

	if i == n {
		return -1
	} else {
		return i
	}
}

func BinarySearch(a []int, n int, x int) int {
	var left int = 0
	var right int = n - 1
	var mid int

	for left <= right {
		mid = (right + left) / 2
		if x == a[mid] {
			return mid
		} else {
			if x < a[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
