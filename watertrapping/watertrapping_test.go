package main

import (
	"strconv"
	"testing"
)

func Test_main(t *testing.T) {
	var field = []int{0, 2, 0, 1, 0, 3}
	var r = solve2(field)
	if r != 5 {
		t.Errorf("result must be 5 is " + strconv.Itoa(r))
	}

}

func Test_main2(t *testing.T) {
	var field = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	var r = solve2(field)
	if r != 6 {
		t.Errorf("result must be 6 is " + strconv.Itoa(r))
	}

}

var maxMap = make(map[int]int)
var maxrMap = make(map[int]int)

func Test_maxLeft(t *testing.T) {
	var field = []int{0, 3, 4, 0, 3}
	var r = maxLeft(4, &field, maxMap)
	if r != 4 {
		t.Errorf("result musst be 4 and is " + strconv.Itoa(r))
	}
}
func Test_maxRight(t *testing.T) {
	var field = []int{0, 3, 4, 0, 3}
	var r = maxRight(2, &field, maxrMap)
	if r != 3 {
		t.Errorf("result musst be 3 and is " + strconv.Itoa(r))
	}
}
func Test_solve(t *testing.T) {
	var field = []int{2, 1, 0, 3}
	var r = solve2(field)
	if r != 3 {
		t.Errorf("result must be 3 is " + strconv.Itoa(r))
	}

}
