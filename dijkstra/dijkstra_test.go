package dijkstra

import (
	"testing"
)

func TestCase1(t *testing.T) {
	graph := make([][]int, 6)
	graph[0] = append(graph[0], 0, 1, 3)
	graph[1] = append(graph[1], 0, 2, 10)
	graph[2] = append(graph[2], 1, 3, 5)
	graph[3] = append(graph[3], 3, 2, 4)
	graph[4] = append(graph[4], 3, 0, 7)
	graph[5] = append(graph[5], 2, 3, 6)

	res := dik(0, 4, graph)
  if res != 10 {
    t.Errorf("must be 10")
  }
}
func TestCase2(t *testing.T) {
	graph := make([][]int, 7)
	graph[0] = append(graph[0], 0, 1, 5)
	graph[1] = append(graph[1], 0, 4, 7)
	graph[2] = append(graph[2], 1, 2, 3)
	graph[3] = append(graph[3], 2, 5, 2)
	graph[4] = append(graph[4], 5, 3, 9)
	graph[5] = append(graph[5], 4, 5, 1)
	graph[6] = append(graph[6], 4, 3, 2)

	res := dik(0, 6, graph)
  if res != 9 {
    t.Errorf("must be 9, is %d", res)
  }
}
func TestCase3(t *testing.T) {
	graph := make([][]int, 4)
	graph[0] = append(graph[0], 0, 1, 3)
	graph[1] = append(graph[1], 1, 3, 2)
	graph[2] = append(graph[2], 3, 2, 0)
	graph[3] = append(graph[3], 2, 3, 1)

	res := dik(0, 4, graph)
  if res != 5 {
    t.Errorf("must be 5 ")
  }
}
func TestCase4(t *testing.T) {
	graph := make([][]int, 5)
	graph[0] = append(graph[0], 0, 1, 3)
	graph[1] = append(graph[1], 1, 3, 2)
	graph[2] = append(graph[2], 3, 2, 0)
	graph[3] = append(graph[3], 2, 3, 1)
	graph[4] = append(graph[4], 0, 2, 1)

	res := dik(0, 4, graph)
  if res != 3 {
    t.Errorf("must be 3")
  }
}
func TestCase5(t *testing.T) {
	graph := make([][]int, 10)
	graph[0] = append(graph[0], 0, 1, 5)
	graph[1] = append(graph[1], 0, 5, 3)
	graph[2] = append(graph[2], 1, 2, 2)
	graph[3] = append(graph[3], 2, 3, 4)
	graph[4] = append(graph[4], 2, 4, 3)
	graph[5] = append(graph[5], 3, 6, 1)
	graph[6] = append(graph[6], 4, 3, 2)
	graph[7] = append(graph[7], 5, 7, 2)
	graph[8] = append(graph[8], 5, 6, 15)
	graph[9] = append(graph[9], 6, 7, 8)

	res := dik(0, 8, graph)
  if res != 12 {
    t.Errorf("must be 12")
  }
}
