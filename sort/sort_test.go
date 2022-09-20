package mysort

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	var s []int
	s = append(s, 1, 4, 55, 3, 2)
	x := selectionSort(s)

	if x[0] != 1 {
		t.Errorf("pos 0 must be 1")
	}
	if x[1] != 2 {
		t.Errorf("pos 1 must be 2")
	}
	if x[2] != 3 {
		t.Errorf("pos 2 must be 3")
	}
	if x[3] != 4 {
		t.Errorf("pos 3 must be 4")
	}
	if x[4] != 55 {
		t.Errorf("pos 4 must be 55")
	}
}
func TestMain2(t *testing.T) {
	var s []int
	s = append(s, 99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0)
	x := selectionSort(s)
	for i := 0; i < len(x); i++ {
		fmt.Println(s[i])
	}
	if x[0] != 0 {
		t.Errorf("pos 0 must be 0")
	}

	if x[3] != 4 {
		t.Errorf("pos 3 must be 4")
	}
	if x[10] != 283 {
		t.Errorf("pos 10 must be 283")
	}
}
func TestMainBubble1(t *testing.T) {
	var s []int
	s = append(s, 1, 4, 55, 3, 2)
	x := bubbleSort(s)

	if x[0] != 1 {
		t.Errorf("pos 0 must be 1")
	}
	if x[1] != 2 {
		t.Errorf("pos 1 must be 2")
	}
	if x[2] != 3 {
		t.Errorf("pos 2 must be 3")
	}
	if x[3] != 4 {
		t.Errorf("pos 3 must be 4")
	}
	if x[4] != 55 {
		t.Errorf("pos 4 must be 55")
	}
}
func TestBubble2(t *testing.T) {
	var s []int
	s = append(s, 99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0)
	x := bubbleSort(s)
	for i := 0; i < len(x); i++ {
		fmt.Println(s[i])
	}
	if x[0] != 0 {
		t.Errorf("pos 0 must be 0")
	}

	if x[3] != 4 {
		t.Errorf("pos 3 must be 4")
	}
	if x[10] != 283 {
		t.Errorf("pos 10 must be 283")
	}
}
