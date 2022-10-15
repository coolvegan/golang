package priority

import "math"

type priority_queue struct {
	heap []int
}

func (p *priority_queue) size() int {
	return len(p.heap)
}

func (p *priority_queue) isEmpty() bool {
	return len(p.heap) == 0
}

func (p *priority_queue) push(value int) {
  p.heap = append(p.heap, value)
}

func (p *priority_queue) siftUp(){
  nodeIdx := len(p.heap)-1
  for nodeIdx > 0 && p.compare(nodeIdx, p.parent(nodeIdx)){
    p.swap(nodeIdx, p.parent(nodeIdx))
    nodeIdx = p.parent(nodeIdx)
  }
}
func (p *priority_queue) siftDown(){
  nodeIdx := 0
  for p.leftChild(nodeIdx) < len(p.size()) && p.compare(p.leftChild(nodeIdx), nodeIdx)
  }
}

func (p *priority_queue) pop() {
  if len(p.heap) > 1 {
    p.swap(0,len(p.heap)-1)
  }
  poppedValue := p.heap[len(p.heap)-1]
  p.siftDown()
  return poppedValue
}

func (p *priority_queue) peek() int {
	return p.heap[0]
}

func (p *priority_queue) parent(idx int) int {
  return int(math.Floor(float64((idx-1)/2)))
}

func (p *priority_queue) leftChild(idx int) int {
  return int(math.Floor(float64((idx*1+1))))
}

func (p *priority_queue) rightChild(idx int) int {
  return int(math.Floor(float64((idx*1+2))))
}

func (p *priority_queue) swap(i,j int){
  tmp := p.heap[i]
  p.heap[i] = p.heap[j]
  p.heap[j] = tmp
}
func (p *priority_queue) compare(i,j int) bool {
  if p.heap[i] > p.heap[j]{
    return true
  }
  return false
}


