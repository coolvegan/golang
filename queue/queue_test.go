package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	var q Queue

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	v, _ := q.dequeue()
	v2, _ := q.dequeue()

	if v != 1 {
		t.Errorf("must be one")
	}
	if v2 != 2 {
		t.Errorf("must be two")
	}
	v3, _ := q.dequeue()
	if v3 != 3 {
		t.Errorf("must be 3")
	}
	v4, _ := q.dequeue()
	if v4 != -1 {
		t.Errorf("must be -1")
	}
}
func TestCaseInitQueue(t *testing.T) {
	var q Queue

	r, b := q.dequeue()
	if r != -1 {
		t.Errorf("r must be -1 ")
	}
	if b != false {
		t.Errorf("b must be false")
	}
}
func TestCase2(t *testing.T) {
	var q Queue

	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)

	v, _ := q.dequeue()
	v2, _ := q.dequeue()

	if v != 1 {
		t.Errorf("must be one")
	}
	if v2 != 2 {
		t.Errorf("must be two")
	}
	v3, _ := q.dequeue()
	if v3 != 3 {
		t.Errorf("must be 3")
	}
	v4, _ := q.dequeue()
	if v4 != -1 {
		t.Errorf("must be -1")
	}
	q.enqueue(4)
	q.enqueue(5)
	q.enqueue(6)
	v, _ = q.dequeue()
	v2, _ = q.dequeue()

	if v != 4 {
		t.Errorf("must be four")
	}
	if v2 != 5 {
		t.Errorf("must be five")
	}
	q.enqueue(7)
	q.enqueue(8)
	v, _ = q.dequeue()
	v2, _ = q.dequeue()
	v3, _ = q.dequeue()

	if v != 6 {
		t.Errorf("must be six")
	}
	if v2 != 7 {
		t.Errorf("must be seven")
	}
	if v3 != 8 {
		t.Errorf("must be eight")
	}
}
