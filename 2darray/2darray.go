package main

type point struct {
	row int
	col int
}

var direction [4]point = [4]point{
	{row: -1, col: 0},
	{row: 0, col: 1},
	{row: 1, col: 0},
	{row: 0, col: -1}}

func dfs(matrix [][]int, row, col int, seen *[][]bool, values *[]int) {
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || (*seen)[row][col] == true {
		return
	}

	*values = append(*values, matrix[row][col])
	(*seen)[row][col] = true

	for i := 0; i < len(direction); i++ {
		currentDir := direction[i]
		dfs(matrix, row+currentDir.row, col+currentDir.col, seen, values)
	}
}

func traversalDFS(matrix *[][]int) []int {
	var result []int
	colCount := len((*matrix)[0])
	rowCount := len(*matrix)
	seen := make([][]bool, rowCount)
	for i := 0; i < colCount; i++ {
		seen[i] = make([]bool, colCount)
	}
	dfs((*matrix), 0, 0, &seen, &result)
	return result
}

func bfs(matrix[][]int, row, col int, seen*[][]bool, values *[]int, queue *[]point){
	if row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0]) || (*seen)[row][col] == true {
		return
	}
  element := matrix[row][col]
  *values = append(*values, element)
  (*seen)[row][col] = true
  var p1 point
  p1.row = row
  p1.col = col
  *queue = append((*queue), p1)

  for i:=0; i < len(direction);i++{
    currentDir := direction[i]
    var px point
    px.row = p1.row + currentDir.row
    px.col = p1.col + currentDir.col
    *queue = append((*queue), px)
  }

  for len(*queue) > 0 {
    mypoint := (*queue)[0]
    *queue = (*queue)[1:]
    bfs(matrix, mypoint.row, mypoint.col, seen, values, queue)
  }
}


func traversalBFS(matrix *[][]int) []int {
  var result[]int
  var queue[]int
	colCount := len((*matrix)[0])
	rowCount := len(*matrix)
	seen := make([][]bool, rowCount)
	for i := 0; i < colCount; i++ {
		seen[i] = make([]bool, colCount)
	}


  bfs(*matrix, 0,0, &seen, &result, &queue)

  return result
}
