package coursescheduler

import (
	"testing"
)

func TestCase1AdjaList(t *testing.T) {
	var matrix [][]int
	matrix = make([][]int, 7)

	matrix[0] = append(matrix[0], 1, 0)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 2, 5)
	matrix[3] = append(matrix[3], 0, 3)
	matrix[4] = append(matrix[4], 4, 3)
	matrix[5] = append(matrix[5], 3, 5)
	matrix[6] = append(matrix[6], 4, 5)
	res := getAdjacentList(matrix)
	if len(res[0]) != 1 {
		t.Errorf("must be 1")
	}
	if len(res[1]) != 1 {
		t.Errorf("must be 1")
	}
	if len(res[2]) != 0 {
		t.Errorf("must be 0")
	}
	if len(res[3]) != 2 {
		t.Errorf("must be 2")
	}
	if len(res[4]) != 0 {
		t.Errorf("must be 0")
	}
	if len(res[5]) != 3 {
		t.Errorf("must be 3")
	}

}
func TestCase2(t *testing.T) {

	matrix := make([][]int, 4)

	matrix[0] = append(matrix[0], 2, 1)
	matrix[1] = append(matrix[1], 0, 2)
	matrix[2] = append(matrix[2], 1, 0)
	matrix[3] = append(matrix[3], 3, 0)
	res := hasCycle(matrix)
	if res != true {
		t.Errorf("must be true")
	}

}
func TestCase3(t *testing.T) {

	matrix := make([][]int, 3)

	matrix[0] = append(matrix[0], 0, 1)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 3, 2)
	res := hasCycle(matrix)
	if res != false {
		t.Errorf("must be false")
	}

}
func TestCase4(t *testing.T) {

	matrix := make([][]int, 4)

	matrix[0] = append(matrix[0], 1, 0)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 3, 2)
	matrix[3] = append(matrix[3], 1, 3)
	res := hasCycle(matrix)
	if res != true {
		t.Errorf("must be true")
	}
}
func TestCase5(t *testing.T) {
	matrix := make([][]int, 7)
	matrix[0] = append(matrix[0], 1, 0)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 2, 5)
	matrix[3] = append(matrix[3], 0, 3)
	matrix[4] = append(matrix[4], 4, 3)
	matrix[5] = append(matrix[5], 3, 5)
	matrix[6] = append(matrix[6], 4, 5)
}

func TestCase6(t *testing.T) {
	matrix := make([][]int, 7)
	matrix[0] = append(matrix[0], 1, 0)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 2, 5)
	matrix[3] = append(matrix[3], 0, 3)
	matrix[4] = append(matrix[4], 4, 3)
	matrix[5] = append(matrix[5], 3, 5)
	matrix[6] = append(matrix[6], 4, 5)
	res := canFinish(5, matrix)
	if res == true {
		t.Errorf("Must be true")
	}

}
func TestCase7(t *testing.T) {

	matrix := make([][]int, 4)
	matrix[0] = append(matrix[0], 1, 0)
	matrix[1] = append(matrix[1], 2, 1)
	matrix[2] = append(matrix[2], 3, 2)
	matrix[3] = append(matrix[3], 1, 3)
	res := canFinish(4, matrix)
	if res == true {
		t.Errorf("must be false")
	}
}
