package sort

func SelectionSort(a []int, n int) {
	var min int

	for i := 0; i < n-1; i++ {
		min = i
		for j := i + 1; j < n; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func InsertionSort(a []int, n int) {
	var pos, i, x int
	for i = 1; i < n; i++ {
		x = a[i]
		pos = i - 1

		for pos >= 0 && a[pos] > x {
			a[pos+1] = a[pos]
			pos--
		}
		a[pos+1] = x
	}
}

func BinaryInsertionSort(a []int, n int) {
	var i, j, x, left, right, mid int

	for i = 1; i < n; i++ {
		x = a[i]
		left = 0
		right = i - 1
		for left <= right {
			mid = (left + right) / 2
			if x < a[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		for j = i - 1; j >= left; j-- {
			a[j+1] = a[j]
		}
		a[left] = x
	}
}

func InterchangeSort(a []int, n int) {
	var i, j int

	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func BubbleSort(a []int, n int) {
	var i, j int

	for i = 0; i < n-1; i++ {
		for j = n - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}

func ShakerSort(a []int, n int) {
	var i, left, right, k int
	left = 0
	right = n - 1

	for left < right {
		for i = right; i > left; i-- {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
				k = i
			}
		}
		left = k

		for i = left; i < right; i++ {
			if a[i+1] < a[i] {
				a[i+1], a[i] = a[i], a[i+1]
				k = i
			}
		}
		right = k
	}
}

//Implement heap sort algorithm
func Shift(a []int, left int, right int) {
	var i, j, x int
	var cont bool = true
	i = left
	j = 2 * i
	x = a[i]
	for j < right-1 && cont == true {
		if a[j] < a[j+1] {
			j = j + 1
		}

		if a[j] <= x {
			cont = false
		} else {
			a[i] = a[j]
			i = j
			j = 2 * i
			a[i] = x
		}
	}
}

func CreateHeap(a []int, n int) {
	var left int = n / 2
	for left >= 0 { //attention here
		Shift(a, left, n)
		left--
	}
}

func HeapSort(a []int, n int) {
	var r int = n - 1 //index of last element in array
	CreateHeap(a, n)
	for r > 0 { //attention here
		a[r], a[0] = a[0], a[r]
		Shift(a, 0, r)
		r--
	}
}

//End Heap sort

//Implement shell sort
func ShellSort(a []int, n int, h []int) {
	var k int = len(h)
	var i, j, step, x, pos int

	for i = 0; i < k; i++ {
		step = h[i]
		for j = step; j < n; j++ {
			x = a[j]
			pos = j - step
			for pos >= 0 && a[pos] > x {
				a[pos+step] = a[pos]
				pos -= step
			}
			a[pos+step] = x
		}
	}
}

func QuickSort(a []int, l int, r int) {
	var i, j, x int
	x = a[(l+r)/2]
	i = l
	j = r

	for i <= j {
		for a[i] < x {
			i++
		}

		for a[j] > x {
			j--
		}

		if i <= j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}

	if l < j {
		QuickSort(a, l, j)
	}

	if i < r {
		QuickSort(a, i, r)
	}
}
