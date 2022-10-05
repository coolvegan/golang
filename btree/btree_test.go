package btree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	var a, b, c, d, e Node

	a.left = &b
	b.left = &c
	a.right = &d
	d.right = &e

	r := treeDepth(&a, 0)

	if r != 3 {
		t.Errorf("Must be 3")
	}

}
func TestCase2(t *testing.T) {
	var a, b Node

	a.left = &b

	r := treeDepth(&a, 0)

	if r != 2 {
		t.Errorf("Must be 2")
	}

}
func TestCase3(t *testing.T) {

	r := treeDepth(nil, 0)

	if r != 0 {
		t.Errorf("Must be 0")
	}

}
func TestCase4(t *testing.T) {
	var a, b, c, d, e Node
	a.left = &b
	b.left = &c
	c.left = &d
	d.left = &e
	r := treeDepth(&a, 0)

	if r != 5 {
		t.Errorf("Must be 5")
	}

}
func TestCase5(t *testing.T) {

	var a, b, c, d, e, f, g Node
	a.left = &b
	b.left = &c
	c.left = &d
	d.left = &e
	e.left = &f
	f.left = &g
	r := treeDepth(&a, 0)
	if r != 7 {
		t.Errorf("Must be 7")
	}
}

func TestQueue1(t *testing.T) {
	var q Queue
  var x,y SingleNode
  x.val = 2
  y.val = 3
	q.Enque(x)
	q.Enque(y)

	if q.head.val != 2 {
		t.Errorf("Must be 2")
	}
	if q.head.next.prev.val != 2 {
		t.Errorf("Must be 2")
	}
	if q.head.next.val != 3 {
		t.Errorf("Must be 3")
	}
}

func TestQueue2(t *testing.T) {
	var q Queue
  var a,y SingleNode
  a.val = 2
  y.val = 3
	q.Enque(a)
	q.Enque(y)
	x, _ := q.Deque()
	if x.val != 3 {
		t.Errorf("Must be 3")
	}
	if q.tail.val != 2 {
		t.Errorf("must be 2")
	}
	x, _ = q.Deque()
  if x.val != 2 {
    t.Errorf("must be 2")
  }
	x, _ = q.Deque()
  if x.val != -1{
    t.Errorf("must be -1")
  }
}

func TestQueue(t *testing.T) {

	var a, b, c, d, e SingleNode
	a.next = &b
	b.next = &c
	c.next = &d
	d.next = &e
  bfs(&a) 
}
