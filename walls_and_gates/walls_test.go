package wallsandgates

import (
	"fmt"
	"math"
	"testing"
)

func TestCase1(t *testing.T) {
	grid := make([][]int, 2)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 2)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][0] = 0

	r := solve(grid)

	if r[0][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][0] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][1] != 2 {
		t.Errorf("Must be 2")
	}
}
func TestCase2(t *testing.T) {
	grid := make([][]int, 3)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 3)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][2] = 0

	r := solve(grid)

	if r[0][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][0] != 3 {
		t.Errorf("Must be 3")
	}
	if r[1][1] != 2 {
		t.Errorf("Must be 2")
	}
	if r[2][0] != 4 {
		t.Errorf("Must be 4")
	}
}
func TestCase3(t *testing.T) {
	grid := make([][]int, 3)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 3)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][2] = 0
	grid[1][1] = -1
	grid[2][2] = -1

	r := solve(grid)

	if r[0][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][0] != 3 {
		t.Errorf("Must be 3")
	}
	if r[2][0] != 4 {
		t.Errorf("Must be 4")
	}
	if r[2][1] != 5 {
		t.Errorf("Must be 5")
	}

}
func TestCase4(t *testing.T) {
	grid := make([][]int, 3)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 3)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][2] = 0
	grid[2][0] = 0
	grid[1][1] = -1
	grid[2][2] = -1

	r := solve(grid)

	if r[0][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][0] != 1 {
		t.Errorf("Must be 1")
	}
	if r[2][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[0][0] != 2 {
		t.Errorf("Must be 2")
	}

}
func TestCase5(t *testing.T) {
	grid := make([][]int, 4)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 4)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][0] = 0
	grid[3][2] = 0

	r := solve(grid)

	if r[0][1] != 1 {
		t.Errorf("Must be 1")
	}
	if r[1][0] != 1 {
		t.Errorf("Must be 1")
	}
	if r[2][1] != 2 {
		t.Errorf("Must be 2")
	}
	if r[0][3] != 3 {
		t.Errorf("Must be 2")
	}

}
func TestCase6(t *testing.T) {
	grid := make([][]int, 4)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 4)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j] = int(math.Inf(1))
		}
	}

	grid[0][0] = 0
	grid[3][2] = 0
	grid[2][2] = -1
	grid[1][2] = -1
	grid[0][1] = -1

	r := solve(grid)

	if r[1][1] != 2 {
		t.Errorf("Must be 2")
	}
	if r[1][0] != 1 {
		t.Errorf("Must be 1")
	}
	if r[2][1] != 2 {
		t.Errorf("Must be 2")
	}
	if r[0][3] != 4 {
		t.Errorf("Must be 4")
	}
	if r[0][2] != 5 {
		t.Errorf("Must be 5")
	}
}

