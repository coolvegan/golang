package rottingoranges

type Point struct {
	row int
	col int
}

func getRotten(grid [][]int) int {
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
	enjoyable := false
	var queue []Point
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				var p Point
				p.row = i
				p.col = j
				queue = append(queue, p)
			}
			if grid[i][j] == 1 {
				enjoyable = true
			}
		}
	}
	if len(queue) == 0 && !enjoyable {
		return 0
	}

	var nextqueue []Point
	count := -1
	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		for i := 0; i < len(directions); i++ {
			chkrow := directions[i][0] + element.row
			chkcol := directions[i][1] + element.col
			if chkrow < 0 || chkrow >= len(grid) || chkcol < 0 || chkcol >= len(grid[0]) || grid[chkrow][chkcol] == 2 {
				continue
			}
			if grid[chkrow][chkcol] == 1 {
				grid[chkrow][chkcol] = 2
				var np Point
				np.row = chkrow
				np.col = chkcol
				nextqueue = append(nextqueue, np)
			}
		}
		if len(queue) == 0 {
			for len(nextqueue) > 0 {
				queue = append(queue, nextqueue[0])
				nextqueue = nextqueue[1:]
			}
			count++
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return count
}
