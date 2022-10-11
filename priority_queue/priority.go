package priority 

type priority_queue struct{
  heap []int
}

func (p* priority_queue) size() int{
  return len(p.heap)
}
