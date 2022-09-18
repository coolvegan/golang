package main

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Peek() (int, bool) {
	return (*s)[len(*s)-1], s.isEmpty() != true
}

func (s *Stack) Pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	}
	val := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return val, true

}
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

type Queue struct {
	s1 Stack
	s2 Stack
}

func (q *Queue) enqueue(val int) {
	q.s1.Push(val)
}
func (q *Queue) dequeue() (int, bool) {
	if q.s2.isEmpty() {
		for !q.s1.isEmpty() {
			v, _ := q.s1.Pop()
			q.s2.Push(v)
		}
	}
	if q.s2.isEmpty() {
		return -1, false
	}
	r, _ := q.s2.Pop()
	return r, true
}
