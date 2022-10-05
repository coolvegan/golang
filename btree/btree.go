package btree

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

type SingleNode struct {
	val  int
	next *SingleNode
	prev *SingleNode
}
type MyNode struct {
	val  int
	left *MyNode
	right *MyNode
}

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func treeDepth(node *Node, count int) int {
	if node == nil {
		return count
	}
	count++
	return max(treeDepth(node.left, count), treeDepth(node.right, count))
}

type Queue struct {
	head *MyNode
	tail *MyNode
}

func (q *Queue) Enque(a MyNode) {
	var s MyNode
	s = a
	if q.head == nil && q.tail == nil {
		q.head = &s
		q.tail = &s
	} else {
		q.tail.left = &s
		s.left = q.tail
		q.tail = q.tail.left
	}
}

func (q *Queue) Deque() (*MyNode, bool) {
	var fr MyNode
	fr.val = -1
	if q.head == nil {
		return &fr, false
	}
	if q.head == q.tail {
		s := q.head
		q.head = nil
		return s, true
	}
	s := q.tail
	p := q.tail.left
	p.left = nil
	q.tail = p
	s = nil
	return s, true
}

func bfs(node *MyNode) {

  var q Queue
  current := node
  for(current.val != -1){
    fmt.Println(current.val)
    q.Enque(*node)

  }
}
