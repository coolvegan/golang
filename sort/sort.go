package mysort

func selectionSort(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
			}
		}
	}
	return data
}
func bubbleSort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1; j++ {
			if data[j] > data[j+1] {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
	return data
}

func selectionSort2(data []int) []int {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
			}
		}
	}
	return data
}
