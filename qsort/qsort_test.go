package main

import (
	"strconv"
	"testing"
)

func TestQSortCase1(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 1, 5, 3, 7, 94, 33, 45)
	dataset2 = append(dataset2, 1, 3, 5, 7, 33, 45, 94)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}
func TestQSortCase2(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 900, 5, 3, 7, 94, 3, 45, -3)
	dataset2 = append(dataset2, -3, 3, 3, 5, 7, 45, 94, 900)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}
func TestQSortCase3(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 7,5,3,2,1,0)
	dataset2 = append(dataset2, 0,1,2,3,5,7)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}
