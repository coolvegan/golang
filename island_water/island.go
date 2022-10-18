package islandwater

type Point struct {
	row int
	col int
}

type IslandNode struct {
	values *[]Point
	next   *IslandNode
}

type IslandList struct {
	head *IslandNode
	tail *IslandNode
}

func (i *IslandList) append(values *[]Point) {
	if i.head == nil {
		var node IslandNode
		node.values = values
		i.head = &node
		i.tail = i.head
	} else {
		var node IslandNode
		node.values = values
		i.tail.next = &node
		i.tail = i.tail.next
	}
}

func (i *IslandList) count() int {
	if i.head == nil {
		return 0
	}
	count := 0
	current := i.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func travelBFS(matrix [][]int, seen *[][]bool, startrow, startcol int) []Point {
	rowcount := len(matrix)
	colcount := len(matrix[0])
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
	var originPoint Point
	originPoint.row = startrow
	originPoint.col = startcol
	queue = append(queue, originPoint)

	var result []Point

	if matrix[originPoint.row][originPoint.col] != 0 {
		result = append(result, originPoint)
		(*seen)[startrow][startcol] = true
	}

	for len(queue) > 0 {
		firstElement := queue[0]
		queue = queue[1:]
		for i := 0; i < len(directions); i++ {
			row := firstElement.row + directions[i][0]
			col := firstElement.col + directions[i][1]
			if row < 0 || col < 0 || row >= rowcount || col >= colcount || (*seen)[row][col] == true || matrix[row][col] == 0 {
				continue
			}
			var newPoint Point
			newPoint.row = row
			newPoint.col = col
			(*seen)[row][col] = true
			queue = append(queue, newPoint)
			result = append(result, newPoint)
		}
	}
	return result
}

func GetIslands(matrix [][]int) IslandList {
	var islandList IslandList
	var seen [][]bool
	seen = make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		seen[i] = make([]bool, len(matrix[0]))
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if (matrix)[row][col] != 0 && seen[row][col] != true {
				r := travelBFS(matrix, &seen, row, col)
				islandList.append(&r)
			}
		}
	}
	return islandList
}
