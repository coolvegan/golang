package main

import "testing"

func TestBrackets(t *testing.T) {
	s := "()"
	r := brackets(&s)
	if r != true {
		t.Errorf("must be true")
	}
}
func TestBrackets2(t *testing.T) {
	s := "([])"
	r := brackets(&s)
	if r != true {
		t.Errorf("must be true")
	}
}
func TestBrackets3(t *testing.T) {
	s := "([])[[]]"
	r := brackets(&s)
	if r != true {
		t.Errorf("must be true")
	}
}
func TestBrackets4(t *testing.T) {
	s := "{{}([])[[]]}"
	r := brackets(&s)
	if r != true {
		t.Errorf("must be true")
	}
}
func TestBrackets5(t *testing.T) {
	s := "{{}([((()))])[[]]}"
	r := brackets(&s)
	if r != true {
		t.Errorf("must be true")
	}
}
func TestStack(t *testing.T) {
	var s Stack
	if s.isEmpty() == false {
		t.Errorf("Stack must be empty")
	}

	s.Push(".")
	if s.isEmpty() == true {
		t.Errorf("Stack must not be empty")
	}
	chr, _ := s.Pop()
	if chr != "." {
		t.Errorf("chr must be .")
	}
	if s.isEmpty() == false {
		t.Errorf("Stack must be empty")
	}

	s.Push("a").Push("b")
	chr, _ = s.Pop()
	if chr != "b" {
		t.Errorf("chr must be b, is " + chr)
	}
	chr, _ = s.Pop()
	if chr != "a" {
		t.Errorf("chr must be b, is " + chr)
	}
	if s.isEmpty() != true {
		t.Errorf("Stack must be empty")
	}
}
