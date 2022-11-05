package graphbfs

func traversalBFS(graph [][]int) []int{
  var queue []int
  var result []int
  seen := map[int]bool{}
  queue = append(queue, 0)

  for len(queue) > 0 {
    vertex := queue[0]
    queue = queue[1:]
    result = append(result, vertex)
    seen[vertex] = true
    connections := graph[vertex]

    for i := 0; i < len(connections); i++ {
      connection := connections[i]
      if !seen[connection] {
        queue = append(queue, connection)
      }
    }
  }
  return result
}


func traversalDFS(vertex int, graph*[][]int, values* []int, seen* map[int]bool) *[]int{
  *values = append((*values), vertex)
  (*seen)[vertex] = true

  connections := (*graph)[vertex]
  for i := 0; i < len(connections); i++ {
    connection := connections[i]
    if !(*seen)[connection]{
      traversalDFS(connection, graph, values, seen)
    }
  }
  return values
}
