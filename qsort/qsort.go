package main

func qsort(data *[]int, left int, right int) {
	if left >= right {
		return
	}
	divider := divide(data, left, right)
	qsort(data, left, divider-1)
	qsort(data, divider+1, right)
}

func divide(data *[]int, left int, right int) int {
	count := 0
	pivot := (*data)[left]
	for i := left + 1; i <= right; i++ {
		if (*data)[i] <= pivot {
			count++
		}
	}

	pivotIndex := left + count
	swap(data, pivotIndex, left)

	i := left
	j := right

	for i < pivotIndex && j > pivotIndex {
		for (*data)[i] <= pivot {
			i++
		}

		for (*data)[j] > pivot {
			j--
		}

		if i < pivotIndex && j > pivotIndex {
			swap(data, i, j)
			i++
			j--
		}
	}
	return pivotIndex
}

func qsdivide(data *[]int, left int, right int) int {
	pivot := (*data)[right]
	divider := left
	for i := left; i < right; i++ {
		if (*data)[i] <= pivot {
			swap(data, divider, i)
			divider++
		}
	}
	swap(data, divider, right)
	return divider
}

func swap(data *[]int, i int, j int) {
	tmp := (*data)[i]
	(*data)[i] = (*data)[j]
	(*data)[j] = tmp
}

func reverse(data []int) []int {
	p1 := 0
	p2 := len(data) - 1
	for p1 < p2 {
		swap(&data, p1, p2)
		p1++
		p2--
	}
	return data
}

func kThLargestElement(data []int, k int) (int, bool) {
	if k > len(data)-1 {
		return -1, false
	}
	idxToFind := len(data) - k
	return quickselect(&data, 0, len(data)-1, idxToFind)
}

func quickselect(data *[]int, left int, right int, idxToFind int) (int, bool) {
	if left < right {
		divider := divide(data, left, right)
		if divider == idxToFind {
			return (*data)[divider], true
		} else if idxToFind < divider {
			return quickselect(data, left, divider-1, idxToFind)
		} else {
			return quickselect(data, divider+1, right, idxToFind)
		}
	}
	return -1, false
}
