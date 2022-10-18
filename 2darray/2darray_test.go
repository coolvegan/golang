package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	rows := 3
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, 3)
	}

	grid[0][0] = 1
	grid[0][1] = 2
	grid[0][2] = 3
	grid[1][0] = 4
	grid[1][1] = 5
	grid[1][2] = 6
	grid[2][0] = 7
	grid[2][1] = 8
	grid[2][2] = 9


  var values []int
  values = traversalDFS(&grid)
	fmt.Println(values)

	if values[0] != 1 {
		t.Errorf("must be %d", 1)
	}
	if values[1] != 2 {
		t.Errorf("must be %d", 2)
	}
	if values[2] != 3 {
		t.Errorf("must be %d", 3)
	}
	if values[3] != 6 {
		t.Errorf("must be %d", 6)
	}
	if values[4] != 9 {
		t.Errorf("must be %d", 9)
	}
	if values[5] != 8 {
		t.Errorf("must be %d", 8)
	}
	if values[6] != 5 {
		t.Errorf("must be %d", 5)
	}
	if values[7] != 4 {
		t.Errorf("must be %d", 4)
	}
	if values[8] != 7 {
		t.Errorf("must be %d", 7)
	}
}

func TestCase2(t *testing.T){
	rows := 3
	grid := make([][]int, rows)
	for i := 0; i < rows; i++ {
		grid[i] = make([]int, 3)
	}

	grid[0][0] = 1
	grid[0][1] = 2
	grid[0][2] = 3
	grid[1][0] = 4
	grid[1][1] = 5
	grid[1][2] = 6
	grid[2][0] = 7
	grid[2][1] = 8
	grid[2][2] = 9


  var values []int
  values = traversalBFSCourse(&grid)
	fmt.Println(values)
	if values[0] != 1 {
		t.Errorf("must be %d", 1)
	}
	if values[1] != 2 {
		t.Errorf("must be %d", 2)
	}
	if values[2] != 4 {
		t.Errorf("must be %d", 4)
	}
	if values[3] != 3 {
		t.Errorf("must be %d", 3)
	}
	if values[4] != 5 {
		t.Errorf("must be %d", 5)
	}
	if values[5] != 7 {
		t.Errorf("must be %d", 7)
	}
	if values[6] != 6 {
		t.Errorf("must be %d", 6)
	}
	if values[7] != 8 {
		t.Errorf("must be %d", 8)
	}
	if values[8] != 9 {
		t.Errorf("must be %d", 9)
	}
}
