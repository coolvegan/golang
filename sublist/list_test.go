package main

import (
	"fmt"
	"testing"
)

func TestFlattenCase3(t *testing.T) {
	var a, z1, z2 node

	a.value = 1
	a.child = &z1
	z1.value = 4
	z2.value = 5
	z2.head = &z1
	z1.tail = &z2
	y := flatten_list(&a)
	if y.tail.value != 4 {
		t.Errorf("Value must be 4")
	}

	if y.tail.tail.value != 5 {
		t.Errorf("Value must be 5")
	}
	print(&a)
}

func TestFlattenCase1(t *testing.T) {
	var a, b, c, z1, z2 node

	a.value = 1
	a.tail = &b
	b.head = &a
	b.value = 2
	b.tail = &c
	c.head = &b
	c.value = 3

	b.child = &z1
	z1.value = 4
	z2.value = 5
	z2.head = &z1
	z1.tail = &z2

	ff := flatten_list(&a)

	if a.tail.value != 2 {
		t.Errorf("Value must be 2")
	}
	if a.tail.tail.value != 4 {
		t.Errorf("Value must be 4")
	}
	if a.tail.tail.tail.value != 5 {
		t.Errorf("Value must be 5")
	}
	if a.tail.tail.tail.tail.value != 3 {
		t.Errorf("Value must be 3")
	}

	print(ff)
}

func print(n *node) {
	for n != nil {
		n = n.tail

	}

}
func TestFlattenCase2(t *testing.T) {
	var a, b, c, z1, z2 node
	var child_of_z2 node

	child_of_z2.value = 6
	z2_c := node{nil, nil, nil, 7}
	z2_d := node{nil, nil, nil, 8}
	child_of_z2.tail = &z2_c
	z2_c.tail = &z2_d
	z2_c.head = &child_of_z2
	z2_d.head = &z2_c

	a.value = 1
	a.tail = &b
	b.head = &a
	b.value = 2
	b.tail = &c
	c.head = &b
	c.value = 3

	b.child = &z1
	z1.value = 4
	z2.value = 5
	z2.head = &z1
	z1.tail = &z2

	z2.child = &child_of_z2

	flatten_list(&a)

	if a.tail.value != 2 {
		t.Errorf("Value must be 2")
	}
	if a.tail.tail.value != 4 {
		t.Errorf("Value must be 4")
	}
	if a.tail.tail.tail.value != 5 {
		t.Errorf("Value must be 5")
	}
	if a.tail.tail.tail.tail.value != 6 {
		t.Errorf("Value must be 6")
	}

	if a.tail.tail.tail.tail.tail.value != 7 {
		t.Errorf("Value must be 7")
	}
	if a.tail.tail.tail.tail.tail.tail.value != 8 {
		t.Errorf("Value must be 8")
	}
	if a.tail.tail.tail.tail.tail.tail.tail.value != 3 {
		t.Errorf("Value must be 3")
	}
}

func TestFlattenCase4(t *testing.T) {
	var a, b, c node
	a.value = 1
	a.tail = &b
	b.head = &a
	b.value = 2
	c.value = 3
	b.tail = &c
	c.tail = &b

	x := findCycle(&a)
	if x != true {
		t.Errorf("Cycle must be found")
	}
}
func TestFlattenCase5(t *testing.T) {
	var a, b, c, d node
	a.value = 1
	a.tail = &b
	b.value = 2
	c.value = 3
	b.tail = &c
	c.tail = &d
	d.tail = &b

	x := findCycleOptimized(&a)
	if x != true {
		t.Errorf("Cycle must be found")
	}
}
func TestFlattenCase6(t *testing.T) {
	var a, b, c, d node
	a.tail = &b
	b.tail = &c
	c.tail = &d
	d.tail = nil

	x := findCycleOptimized(&a)
	if x != false {
		t.Errorf("Cycle must be false")
	}
}
func TestFlattenCase7(t *testing.T) {
	var a node
	a.tail = nil
	x := findCycleOptimized(&a)
	if x != false {
		t.Errorf("Cycle must be false")
	}
}
func TestFlattenCase8(t *testing.T) {
	var a, b node
	a.tail = &b
	b.tail = &a
	x := findCycleOptimized(&a)
	if x != true {
		t.Errorf("Cycle must be true")
	}
}
func TestFlattenCase9(t *testing.T) {
	var a node
	a.tail = &a
	x := findCycleOptimized(&a)
	if x != true {
		t.Errorf("Cycle must be true")
	}
}
func TestFlattenCaseFindNode1(t *testing.T) {
	var a, b, c, d node
	a.value = 1
	a.tail = &b
	b.value = 2
	c.value = 3
	d.value = 4
	b.tail = &c
	c.tail = &d
	d.tail = &b

	x := findCycleStartingNode(&a)
	if &*x != &b {
		t.Errorf("Cycle must be a")
		fmt.Printf("is : %d", x.value)
	}
}
func TestFlattenCaseFindNode2(t *testing.T) {
	var a, b, c, d, e node
	a.value = 1
	a.tail = &b
	b.value = 2
	c.value = 3
	d.value = 4
	e.value = 5
	b.tail = &c
	c.tail = &d
	d.tail = &e
	e.tail = &c

	x := findCycleStartingNode(&a)
	if &*x != &c {
		t.Errorf("Cycle must be c")
		fmt.Printf("is : %d", x.value)
	}
}
func TestFlattenCaseFindNode3(t *testing.T) {
	var a, b node
	a.tail = &b
	b.tail = &a
	x := findCycleStartingNode(&a)
	if x != &a {
		t.Errorf("Cycle must be true")
	}
}
func TestFlattenCaseFindNode4(t *testing.T) {
	var a node
	a.tail = &a
	x := findCycleStartingNode(&a)
	if x != &a {
		t.Errorf("Cycle must be a")
	}
}
func TestFlattenCaseFindNode5(t *testing.T) {
	var a, b, c node
	a.tail = &b
	a.value = 1
	b.value = 2
	c.value = 3
	b.tail = &c
	c.tail = &c
	x := findCycleStartingNode(&a)
	if x != &c {
		t.Errorf("Cycle must be c")
		fmt.Printf("is : %d", x.value)
	}
}
