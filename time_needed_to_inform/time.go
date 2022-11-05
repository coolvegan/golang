package timeneededtoinform


func getAdjacentList(managers *[]int) [][]int {
  length := len(*managers)
  list := make([][]int, length)
  for i := 0; i < length; i++ {
    manager := (*managers)[i]
    if manager != -1 {
      list[manager] = append(list[manager], i)
    }
  }
  return list
}


func getMaxTime(managers *[]int, informTime *[]int) int{
  var startindex int
  for i := 0; i < len(*managers); i++ {
    if (*managers)[i] == -1 {
      startindex = i
      break
    }
  }
  list := getAdjacentList(managers)
  return dfs(&list, informTime, startindex)
}

func max(v1, v2 int) int{
  if v1 >= v2{
    return v1
  }
  return v2
}


func dfs(list*[][]int, informTime*[]int, index int) int {
  time := (*informTime)[index]
  if time == 0 {
    return 0
  }
  value := 0
  for i := 0; i < len((*list)[index]); i++ { 
    value = max(dfs(list, informTime, (*list)[index][i]), value)
  }
  return time+value
}
