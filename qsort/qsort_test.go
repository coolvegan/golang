package main

import (
	"strconv"
	"testing"
)

func TestQSortCase1(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 1, 5, 3, 7, 94, 33, 45)
	dataset2 = append(dataset2, 1, 3, 5, 7, 33, 45, 94)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}
func TestQSortCase2(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 900, 5, 3, 7, 94, 3, 45, -3)
	dataset2 = append(dataset2, -3, 3, 3, 5, 7, 45, 94, 900)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}
func TestQSortCase3(t *testing.T) {
	var dataset, dataset2 []int
	dataset = append(dataset, 7, 5, 3, 2, 1, 0)
	dataset2 = append(dataset2, 0, 1, 2, 3, 5, 7)

	qsort(&dataset, 0, len(dataset)-1)

	for i := 0; i < len(dataset); i++ {
		if dataset[i] != dataset2[i] {
			t.Errorf("Error: Position " + strconv.Itoa(i) + " must be " + strconv.Itoa(dataset[i]) + " is " + strconv.Itoa(dataset2[i]))
		}
	}
}

func TestKthCase1(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 7, 5, 3, 2, 1, 0)
	k, _ := kThLargestElement(dataset, 2)
	if k != 5 {
		t.Errorf("k must be 5, is " + strconv.Itoa(k))
	}
}
func TestKthCase2(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 7, 5, 3, 2, 1, 0)
	k, _ := kThLargestElement(dataset, 2)
	if k != 5 {
		t.Errorf("k must be 5, is " + strconv.Itoa(k))
	}
}
func TestKthCase3(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 0)
	k, x := kThLargestElement(dataset, 4)
	if x != false {
		t.Errorf("result must be false " + strconv.Itoa(k))
	}
}
func TestKthCase4(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 0, 4, 6)
	k, x := kThLargestElement(dataset, 100)
	if x != false {
		t.Errorf("result must be false " + strconv.Itoa(k))
	}
}

func TestBSearch(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 6)
	r, b := bsearch(&dataset, 0, len(dataset)-1, 2)
	if r != -1 {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	if b != false {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearch2(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 6)
	r, b := bsearch(&dataset, 0, len(dataset)-1, 4)
	if r != 2 {
		t.Errorf(" Error, Result is not in Data, must be 2")
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearch3(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 100, 102, 144)
	r, b := bsearch(&dataset, 0, len(dataset)-1, 144)
	if r != 6 {
		t.Errorf(" Error, Result is not in Data, must be 6")
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearch4(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 100, 102, 144)
	r, b := bsearch(&dataset, 0, len(dataset)-1, 1)
	if r != 0 {
		t.Errorf("Error, Result is not in Data, must be 0")
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearch5(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 100, 102, 144)
	r, b := bsearch(&dataset, 0, len(dataset)-1, 3)
	if r != 1 {
		t.Errorf("Error, Result is not in Data, must be 1, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 4)
	if r != 2 {
		t.Errorf("Error, Result is not in Data, must be 2, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 5)
	if r != 3 {
		t.Errorf("Error, Result is not in Data, must be 3, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 100)
	if r != 4 {
		t.Errorf("Error, Result is not in Data, must be 4, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 102)
	if r != 5 {
		t.Errorf("Error, Result is not in Data, must be 5, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 144)
	if r != 6 {
		t.Errorf("Error, Result is not in Data, must be 6, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearchIterativ1(t *testing.T) {
	var dataset []int
	dataset = append(dataset)
	r, b := bsearchIt(&dataset, 1)
	if r != -1 {
		t.Errorf("Error, Result is not in Data, must be 0")
	}
	if b != false {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearchIterativ2(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 4, 5, 100, 102, 144)
	r, b := bsearchIt(&dataset, 3)
	if r != 1 {
		t.Errorf("Error, Result is not in Data, must be 1, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 4)
	if r != 2 {
		t.Errorf("Error, Result is not in Data, must be 2, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 5)
	if r != 3 {
		t.Errorf("Error, Result is not in Data, must be 3, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 100)
	if r != 4 {
		t.Errorf("Error, Result is not in Data, must be 4, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 102)
	if r != 5 {
		t.Errorf("Error, Result is not in Data, must be 5, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
	r, b = bsearch(&dataset, 0, len(dataset)-1, 144)
	if r != 6 {
		t.Errorf("Error, Result is not in Data, must be 6, is " + strconv.Itoa(r))
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearchIterativ3(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 5, 8, 9, 99, 100)
	r, b := bsearchIt(&dataset, 100)
	if r != 6 {
		t.Errorf("Error, Result is not in Data, must be 6")
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestBSearchIterativ4(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 2, 5, 8, 9, 99, 100)
	r, b := bsearchIt(&dataset, 1)
	if r != 0 {
		t.Errorf("Error, Result is not in Data, must be 0")
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}

func TestFindStartAndIndex(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 2, 3, 3, 5, 8, 9, 99, 100)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 2 || e != 3 {
		t.Errorf("Error, Result is not in Data, must be 2 and 3, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex2(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 3, 3, 3, 8, 9, 99, 100)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 1 || e != 4 {
		t.Errorf("Error, Result is not in Data, must be 1 and 4, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex6(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 3, 3, 3, 3, 3, 8, 9, 99, 100)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 0 || e != 4 {
		t.Errorf("Error, Result is not in Data, must be 0 and 4, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex3(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 2, 3, 3, 3, 8, 9, 99, 100)
	s, e, b := findStartAndEndIndex(&dataset, 2)
	if s != 1 || e != 1 {
		t.Errorf("Error, Result is not in Data, must be 2 and 4, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex4(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 2, 3, 3, 3, 8, 9, 99, 100)
	s, e, b := findStartAndEndIndex(&dataset, 4)
	if s != -1 || e != -1 {
		t.Errorf("Error, Result is not in Data, must be -1 and -1, is %d and %d", s, e)
	}
	if b != false {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex5(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 1, 3, 3, 3, 3, 3, 34)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 1 || e != 5 {
		t.Errorf("Error, Result is not in Data, must be 1 and 5, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex7(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 3, 3, 3, 3, 3, 3, 3)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 0 || e != 6 {
		t.Errorf("Error, Result is not in Data, must be 0 and 6, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
func TestFindStartAndIndex8(t *testing.T) {
	var dataset []int
	dataset = append(dataset, 2, 3, 3, 3, 3, 3, 5)
	s, e, b := findStartAndEndIndex(&dataset, 3)
	if s != 1 || e != 5 {
		t.Errorf("Error, Result is not in Data, must be 1 and 5, is %d and %d", s, e)
	}
	if b != true {
		t.Errorf(" Error, Result is not in Data, must be false")
	}
}
