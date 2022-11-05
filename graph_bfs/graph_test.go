package graphbfs

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	adjacentList := make([][]int, 9)
	adjacentList[0] = append(adjacentList[0], 1, 3)
	adjacentList[1] = append(adjacentList[1], 0)
	adjacentList[2] = append(adjacentList[2], 3, 8)
	adjacentList[3] = append(adjacentList[3], 0, 4, 5, 2)
	adjacentList[4] = append(adjacentList[4], 3, 6)
	adjacentList[5] = append(adjacentList[5], 3)
	adjacentList[6] = append(adjacentList[6], 4, 7)
	adjacentList[7] = append(adjacentList[7], 6)
	adjacentList[8] = append(adjacentList[8], 2)

	res := traversalBFS(adjacentList)
	fmt.Println(res)
}

func TestCase2(t *testing.T) {
	adjacentList := make([][]int, 9)
	adjacentList[0] = append(adjacentList[0], 1, 3)
	adjacentList[1] = append(adjacentList[1], 0)
	adjacentList[2] = append(adjacentList[2], 3, 8)
	adjacentList[3] = append(adjacentList[3], 0, 4, 5, 2)
	adjacentList[4] = append(adjacentList[4], 3, 6)
	adjacentList[5] = append(adjacentList[5], 3)
	adjacentList[6] = append(adjacentList[6], 4, 7)
	adjacentList[7] = append(adjacentList[7], 6)
	adjacentList[8] = append(adjacentList[8], 2)


  seen := map[int]bool{}
  var values []int
	res := traversalDFS(0, &adjacentList, &values, &seen)
	fmt.Println(res)
}
