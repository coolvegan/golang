package islandwater

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	directions := make([][]int, 4)
	directions[0] = make([]int, 2)
	directions[1] = make([]int, 2)
	directions[2] = make([]int, 2)
	directions[3] = make([]int, 2)

	directions[0][0] = -1
	directions[0][1] = 0

	directions[1][0] = 0
	directions[1][1] = 1

	directions[2][0] = 1
	directions[2][1] = 0

	directions[3][0] = 0
	directions[3][1] = -1

	var queue []Point

	var x Point
	x.row = 1
	x.col = 1

	queue = append(queue, x)

	queue = queue[1:]

}

func TestCase2(t *testing.T) {
	matrix := make([][]int, 3)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 4)
	}
	matrix[0][0] = 1
	matrix[0][1] = 1
	matrix[0][2] = 0
	matrix[0][3] = 1
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 0
	matrix[1][3] = 1
	matrix[2][0] = 0
	matrix[2][1] = 1
	matrix[2][2] = 1
	matrix[2][3] = 0

	result := GetIslands(matrix)

	if result.count() != 3 {
		t.Errorf("Must be 3 islands")
	}

}

func TestList1(t *testing.T) {
	var is IslandList
	var p1, p2, p3 Point
	p1.col = 0
	p1.row = 0
	p2.col = 1
	p2.row = 2
	p3.col = 1
	p3.row = 2
	var points, points2, points3 []Point
	points = append(points, p1)
	points2 = append(points2, p2)
	is.append(&points)

	if is.count() != 1 {
		t.Errorf("must be one")
	}
	is.append(&points2)

	if is.count() != 2 {
		t.Errorf("must be two")
	}
	points3 = append(points3, p3)
	is.append(&points3)
	if is.count() != 3 {
		t.Errorf("must be three")
	}
}
func TestCase4(t *testing.T) {
	matrix := make([][]int, 5)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 5)
	}
	matrix[0][0] = 1
	matrix[0][1] = 1
	matrix[0][2] = 0
	matrix[0][3] = 1
	matrix[0][4] = 1
	matrix[1][0] = 0
	matrix[1][1] = 0
	matrix[1][2] = 0
	matrix[1][3] = 0
	matrix[1][4] = 0
	matrix[2][0] = 1
	matrix[2][1] = 1
	matrix[2][2] = 0
	matrix[2][3] = 1
	matrix[2][4] = 1
	matrix[3][0] = 0
	matrix[3][1] = 0
	matrix[3][2] = 0
	matrix[3][3] = 0
	matrix[3][4] = 0
	matrix[4][0] = 1
	matrix[4][1] = 1
	matrix[4][2] = 0
	matrix[4][3] = 1
	matrix[4][4] = 1

	result := GetIslands(matrix)
	current := result.head
	for current != nil {
		fmt.Println(current.values)
		current = current.next
	}

	if result.count() != 6 {
		t.Errorf("Must be 6 islands")
	}

}

func TestCase5(t *testing.T) {
	matrix := make([][]int, 5)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, 5)
	}
	matrix[0][0] = 1
	matrix[0][1] = 0
	matrix[0][2] = 0
	matrix[0][3] = 0
	matrix[0][4] = 1
	matrix[1][0] = 0
	matrix[1][1] = 1
	matrix[1][2] = 1
	matrix[1][3] = 1
	matrix[1][4] = 0
	matrix[2][0] = 0
	matrix[2][1] = 1
	matrix[2][2] = 0
	matrix[2][3] = 0
	matrix[2][4] = 0
	matrix[3][0] = 0
	matrix[3][1] = 1
	matrix[3][2] = 1
	matrix[3][3] = 1
	matrix[3][4] = 0
	matrix[4][0] = 1
	matrix[4][1] = 0
	matrix[4][2] = 0
	matrix[4][3] = 0
	matrix[4][4] = 1

	result := GetIslands(matrix)
	current := result.head
	for current != nil {
		fmt.Println(current.values)
		current = current.next
	}

	if result.count() != 5 {
		t.Errorf("Must be 5 islands")
	}

}

