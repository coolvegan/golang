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

func swap(data *[]int, i int, j int) {
	tmp := (*data)[i]
	(*data)[i] = (*data)[j]
	(*data)[j] = tmp
}
