package wallsandgates

type Point struct {
  row int
  col int
}

func getDirections() [][]int{
  result := make([][]int, 4)
  for direction := 0; direction < len(result); direction++ { 
    result[direction] = make([]int,2)
  }
  //up
  result[0][0] = -1
  result[0][1] = 0
  //right
  result[1][0] = 0
  result[1][1] = 1
  //down
  result[2][0] = 1
  result[2][1] = 0
  //left
  result[3][0] = 0
  result[3][1] = -1
  return result
}

func solve(grid [][]int) [][]int{
  const gate = 0
  const wall = -1
  var queue []Point
  directions := getDirections()
  for row := 0; row < len(grid); row++ {
    for col := 0; col < len(grid[0]); col++ {
      if (grid[row][col] == gate){
        var p Point
        p.row = row
        p.col = col
        queue = append(queue, p)
      }
    }
  }
  for len(queue) > 0 {
    pos := queue[0]
    queue = queue[1:]
    actValue := grid[pos.row][pos.col] + 1
    for direction := 0; direction < len(directions); direction++ {
      var np Point
      np.row = pos.row + directions[direction][0]
      np.col = pos.col + directions[direction][1]
      if np.row < 0 || np.row >= len(grid) || np.col < 0 || np.col >= len(grid[0]) || grid[np.row][np.col] == gate || grid[np.row][np.col] == wall{
        continue
      }
      gridPosValue := grid[np.row][np.col]
      if actValue < gridPosValue{
        grid[np.row][np.col] = actValue
        queue = append(queue, np)
      }
    }
  }

  return grid
}
