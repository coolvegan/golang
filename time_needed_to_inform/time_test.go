package timeneededtoinform

import "testing"

func TestCase1(t *testing.T) {
  var managers, informTime []int
  managers = append(managers, 2,2,4,6,-1,4,4,5)
  informTime = append(informTime, 0,0,4,0,7,3,6,0)
  list := getAdjacentList(&managers)
  if len(list[2]) != 2{
    t.Errorf("must be 2")
  }
  if len(list[4]) != 3{
    t.Errorf("must be 3")
  }
  if len(list[5]) != 1{
    t.Errorf("must be 1")
  }
  if len(list[6]) != 1{
    t.Errorf("must be 1")
  }
  if list[4][2] != 6{
    t.Errorf("must be 6, is %d", list[4][2])
  }
  if list[6][0] != 3{
    t.Errorf("must be 3, is %d", list[6][0])
  }
}
func TestCase2(t *testing.T) {
  var managers, informTime []int
  managers = append(managers, 2,2,4,6,-1,4,4,5)
  informTime = append(informTime, 0,0,4,0,7,3,6,0)
  res := getMaxTime(&managers, &informTime)
  if res != 13 {
    t.Errorf("must be 13")
  }
}
func TestCase3(t *testing.T) {
  var managers, informTime []int
  managers = append(managers, 3,0,0,-1)
  informTime = append(informTime, 5,0,0,10)
  res := getMaxTime(&managers, &informTime)
  if res != 15 {
    t.Errorf("must be 15")
  }
}
func TestCase4(t *testing.T) {
  var managers, informTime []int
  managers = append(managers, 5,-1,5,5,7,4,7,1,1)
  informTime = append(informTime, 0,10,0,0,2,5,0,8,0)
  res := getMaxTime(&managers, &informTime)
  if res != 25 {
    t.Errorf("must be 25, is %d", res)
  }
}

func TestCase5(t *testing.T) {
  var managers, informTime []int
  managers = append(managers, 2,2,-1,0,1,3,3,4,4)
  informTime = append(informTime, 4,2,3,2,3,0,0,0,0)
  res := getMaxTime(&managers, &informTime)
  if res != 9 {
    t.Errorf("must be 9, is %d", res)
  }
}

