package btree2

import (
	"fmt"
	"testing"
)

func TestQueue1(t *testing.T) {
	var q queue

	var n, x, y, z node
	n.val = 1
	x.val = 2
	y.val = 3
	z.val = 4
	q.enque(&n)

	if q.getSize() != 1 {
		t.Errorf("Must be one")
	}
	q.enque(&x)
	q.enque(&y)
	q.enque(&z)
	if q.getSize() != 4 {
		t.Errorf("Must be one")
	}

	r1, _ := q.deque()
	r2, _ := q.deque()
	r3, _ := q.deque()
	r4, _ := q.deque()
	r5, b := q.deque()

	if r1.val != 1 {
		t.Errorf("must be one")
	}
	if r2.val != 2 {
		t.Errorf("must be two")
	}
	if r3.val != 3 {
		t.Errorf("must be three")
	}
	if r4.val != 4 {
		t.Errorf("must be four")
	}
	if q.getSize() != 0 {
		t.Errorf("size must be 0")
	}
	if r5.val != -1 {
		t.Errorf("must be -1")
	}
	if b != false {
		t.Errorf("must be false")
	}
}

func Test2(t *testing.T) {
	var a, b, c, d, e, f, i Treenode
	a.Val = 1
	b.Val = 2
	c.Val = 3
	d.Val = 4
	e.Val = 5
	f.Val = 6
	i.Val = 999
	a.Left = &c
	a.Right = &b
	b.Left = &d
	b.Right = &e
	d.Left = &f
	d.Right = &i
	r := bfs(&a)
	if r[0] != 1 {
		t.Errorf("must be 1")
	}
	if r[1] != 3 {
		t.Errorf("must be 3, is %d", r[1])
	}
	if r[2] != 2 {
		t.Errorf("must be 2, is %d", r[2])
	}
	if r[3] != 4 {
		t.Errorf("must be 4, is %d", r[3])
	}
	if r[4] != 5 {
		t.Errorf("must be 5, is %d", r[4])
	}
	if r[6] != 999 {
		t.Errorf("must be 999, is %d", r[6])
	}
	if r[5] != 6 {
		t.Errorf("must be 6, is %d", r[5])
	}
}

func TestC2(t *testing.T) {
	var a, b, c, d, e, f, h, i, o, p, q, s Treenode
	a.Val = 1
	b.Val = 2
	c.Val = 3
	d.Val = 4
	e.Val = 5
	f.Val = 6
	h.Val = 55
	i.Val = 999
	o.Val = 111
	p.Val = 222
	q.Val = 444
	s.Val = 666
	a.Left = &c
	a.Right = &b
	b.Left = &d
	b.Right = &e
	d.Left = &f
	d.Right = &i
	f.Left = &h
	e.Left = &o
	e.Right = &p
	o.Left = &q
	p.Right = &s
	r := levelOrder(&a)
	fmt.Printf("%d", r)
	if len(r) != 5 {
		t.Errorf("Must be 5, is %d", len(r))
	}
	if len(r[1]) != 2 {
		t.Errorf("Must be 2, is %d", len(r[1]))
	}
	if len(r[3]) != 4 {
		t.Errorf("Must be 4, is %d", len(r[3]))
	}
	if r[3][0] != 6 {
		t.Errorf("Must be 6, is %d", r[3][0])
	}
	if r[3][3] != 222 {
		t.Errorf("Must be 222, is %d", r[3][3])
	}
	if r[4][1] != 444 {
		t.Errorf("Must be 444, is %d", r[4][1])
	}
}


func TestCreateTree1(t *testing.T) {
  var a []int
  a = append(a, 3,9,20,-1,-1,15,7)
  x := createTree(a)
  y := levelOrder(&x)
  fmt.Printf("%d", y)
}
func TestCreateTree3(t *testing.T) {
	var a, b, c Treenode
  a.Val = 1
  b.Val = 2 
  c.Val = 3
  a.Left = &b
  a.Right = &c
  y := RightMaskedlevelOrder(&a)
  if y[1] != 3 {
    t.Errorf("Must be three")
  }
}

func TestRmloDfs(t *testing.T){
  var a,b,c, x, y, z, zz, zzz Treenode
  x.Val = 6
  y.Val = 7
  z.Val = 9
  a.Val = 1
  b.Val = 2
  c.Val = 3
  zz.Val = 999
  zzz.Val = 1000
  a.Right = &b
  a.Left = &x
  b.Right = &c
  b.Left = &y
  y.Left = &z
  c.Right = &zz
  z.Right = &zzz
  var result []int
  RightMaskedLevelOrderDfs(&a, 1, &result)
  if result[0] != 1 {
    t.Errorf("Must be one")
  }
  if result[1] != 2 {
    t.Errorf("Must be two")
  }
  if result[2] != 3 {
    t.Errorf("Must be three")
  }
  if result[3] != 999 {
    t.Errorf("Must be 999")
  }
  if result[4] != 1000 {
    t.Errorf("Must be 1000")
  }
  fmt.Printf("%d",result)

}
func TestRmloBfs(t *testing.T){
  var a,b,c, x, y, z, zz, zzz Treenode
  x.Val = 6
  y.Val = 7
  z.Val = 9
  a.Val = 1
  b.Val = 2
  c.Val = 3
  zz.Val = 999
  zzz.Val = 1000
  a.Right = &b
  a.Left = &x
  b.Right = &c
  b.Left = &y
  y.Left = &z
  c.Right = &zz
  z.Right = &zzz
  var result []int
  result = RightMaskedlevelOrder(&a)
  if result[0] != 1 {
    t.Errorf("Must be one")
  }
  if result[1] != 2 {
    t.Errorf("Must be two")
  }
  if result[2] != 3 {
    t.Errorf("Must be three")
  }
  if result[3] != 999 {
    t.Errorf("Must be 999")
  }
  if result[4] != 1000 {
    t.Errorf("Must be 1000")
  }
  fmt.Printf("%d",result)

}
