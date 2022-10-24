package rottingoranges

import "testing"

func TestCase1(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 0
	matrix[2][0] = 0
	matrix[2][1] = 0
	matrix[2][2] = 0

	r := getRotten(matrix)
	if r != 2 {
		t.Errorf("must be two, is %d", r)
	}
}
func TestCase2(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[1][0] = 0
	matrix[1][1] = 1
	matrix[1][2] = 0
	matrix[2][0] = 0
	matrix[2][1] = 0
	matrix[2][2] = 0

	r := getRotten(matrix)
	if r != 2 {
		t.Errorf("must be two, is %d", r)
	}
}
func TestCase3(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 1
	matrix[2][0] = 0
	matrix[2][1] = 0
	matrix[2][2] = 0

	r := getRotten(matrix)
	if r != 3 {
		t.Errorf("must be three, is %d", r)
	}
}
func TestCase4(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 1
	matrix[2][0] = 1
	matrix[2][1] = 1
	matrix[2][2] = 1

	r := getRotten(matrix)
	if r != 6  {
		t.Errorf("must be six, is %d", r)
	}
}
func TestCase5(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 0
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 1
	matrix[2][0] = 2
	matrix[2][1] = 1
	matrix[2][2] = 1

	r := getRotten(matrix)
	if r != 3  {
		t.Errorf("must be three, is %d", r)
	}
}
func TestCase6(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 3)
	}
	matrix[0][0] = 2
	matrix[0][1] = 1
	matrix[0][2] = 1
	matrix[1][0] = 0
	matrix[1][1] = 1
	matrix[1][2] = 1
	matrix[2][0] = 1
	matrix[2][1] = 0
	matrix[2][2] = 1

	r := getRotten(matrix)
	if r != -1  {
		t.Errorf("must be -1, is %d", r)
	}
}
