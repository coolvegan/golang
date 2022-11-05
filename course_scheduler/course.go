package coursescheduler

func getAdjacentList(matrix [][]int) [][]int {
	length := len(matrix)
	list := make([][]int, length)
	for i := 0; i < length; i++ {
		root := matrix[i][1]
		dependent := matrix[i][0]
		list[root] = append(list[root], dependent)
	}
	return list
}

func hasCycle(matrix [][]int) bool {
	list := getAdjacentList(matrix)
	seen := map[int]bool{}
	return dfs(list, 0, &seen)
}

func dfs(matrix [][]int, index int, seen *map[int]bool) bool {
	if len(matrix[index]) == 0 {
		return false
	}
	if (*seen)[index] {
		return true
	}
	var result = false
	(*seen)[index] = true
	for i := 0; i < len(matrix[index]); i++ {
		result = result || dfs(matrix, matrix[index][i], seen)
	}
	return result
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func getIndegree(matrix [][]int) []int {
	length := len(matrix)
	list := make([]int, length)
	for i := 0; i < length; i++ {
		dep := matrix[i][0]
		list[dep] += 1
	}
	return list
}

func canFinish(n int, matrix [][]int) bool {
	inDegree := make([]int, n)
  for i := 0; i < n; i++ {
    inDegree[i]  = 0
  }
  adjList := make(map[int][]int)

  for i := 0; i < len(matrix); i++ {
    pair := matrix[i]
    inDegree[pair[0]]++
    adjList[pair[1]] = append(adjList[pair[1]],  pair[0])
  }
  var stack []int

  for i := 0; i < len(inDegree); i++ {
    if(inDegree[i] == 0){
      stack = append(stack, i)
    }  
  }

  count := 0

  for len(stack)>0{
    element := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    count++
    adjacent := adjList[element]
    for i := 0; i < len(adjacent); i++ {
      next := adjacent[i] 
      inDegree[next]--
      if(inDegree[next] == 0){
        stack = append(stack,next)
      }
    }
  }
  return count == n
}
