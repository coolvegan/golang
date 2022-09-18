package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.T) {
	var a, b, c node

	a.value = 1
	b.value = 2
	a.next = &b
	b.next = &c
	c.value = 3
	c.next = nil

	x := reverse_list(&a)
	if x.value != 3 {
		m.Errorf("Error must be 3")
	}
	var xx, yy node
	xx.value = 1
	xx.next = &yy
	yy.value = 2

	x = reverse_list(&xx)
	if x.value != 2 {
		m.Errorf("Error must be 2")
		fmt.Printf("is %d", x.value)
	}
}

func TestMain2(m *testing.T) {
	var a, b, c, d, e node

	a.value = 1
	b.value = 2
	a.next = &b
	b.next = &c
	c.value = 3
	c.next = &d
	d.value = 4
	d.next = &e
	e.value = 5
	e.next = nil
	current := &a
	current = reverse_sub_list(current, 2, 3)
	if current.next.next.value != 2 {
		m.Errorf("Must be two")
	}
	if current.next.next.next.value != 4 {
		m.Errorf("Must be four")
	}
	reverse_sub_list(current, 2, 3)
	if current.next.next.value != 3 {
		m.Errorf("Must be three")
	}
	if current.next.next.next.value != 4 {
		m.Errorf("Must be four")
	}

}

func print(current *node) {
	for current != nil {
		fmt.Printf("%d", current.value)
		current = current.next
	}
	fmt.Println()
}

// fmt.Printf("la: %d mi: %d ni: %d lr: %d\n", la.value, mi.value, ni.value, lr.value)
func TestMain3(m *testing.T) {
	var a, b, c, d, e node

	a.value = 1
	b.value = 2
	a.next = &b
	b.next = &c
	c.value = 3
	c.next = &d
	d.value = 4
	d.next = &e
	e.value = 5
	e.next = nil
	current := &a

	current = reverse_sub_list(current, 1, 3)
	if current.next.next.value != 1 {
		m.Errorf("Must be one")
		fmt.Printf("is %d", current.next.next.value)
	}
	if current.next.next.next.value != 4 {
		m.Errorf("Must be four")
	}
	if current.value != 3 {
		m.Errorf("Must be three")
	}
	current = reverse_sub_list(current, 1, 3)
	if current.next.next.value != 3 {
		m.Errorf("Must be three")
	}
	if current.next.next.next.value != 4 {
		m.Errorf("Must be four")
	}

}

func TestMain4(m *testing.T) {
	var a, b, c, d, e node

	a.value = 1
	b.value = 2
	a.next = &b
	b.next = &c
	c.value = 3
	c.next = &d
	d.value = 4
	d.next = &e
	e.value = 5
	e.next = nil
	current := &a
	x := reverse_sub_list(current, 1, 2)
	if x.value != 2 {
		m.Errorf("Must be two")
		fmt.Printf("is %d", x.value)
	}

	if x.next.value != 1 {
		m.Errorf("Must be one")
		fmt.Printf("is %d", current.next.value)
	}

	var single node
	single.value = 22
	single.next = nil

	r := reverse_sub_list(&single, 1, 1)
	if r.value != 22 {
		m.Errorf("Must be 22")
		fmt.Printf("is %d", r.value)
	}

}
