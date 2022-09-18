package main

import "testing"

func TestMain(t *testing.T) {
	s := "a(b)"
	r := MinRemoveToMakeValid(s)
	if r != "a(b)" {
		t.Errorf("must be a(b)")
	}
}
func TestMain2(t *testing.T) {
	s := ")a(b)"
	r := MinRemoveToMakeValid(s)
	if r != "a(b)" {
		t.Errorf("must be a(b), is " + r)
	}
}
func TestMain3(t *testing.T) {
	s := ")a))(b)"
	r := MinRemoveToMakeValid(s)
	if r != "a(b)" {
		t.Errorf("must be a(b), is " + r)
	}
}
func TestMain4(t *testing.T) {
	s := ")a())(b)"
	r := MinRemoveToMakeValid(s)
	if r != "a()(b)" {
		t.Errorf("must be a(b), is " + r)
	}
}
func TestMain5(t *testing.T) {
	s := ")))((("
	r := MinRemoveToMakeValid(s)
	if r != "" {
		t.Errorf("must be a(b), is " + r)
	}
}
func TestMain6(t *testing.T) {
	s := "ab(cd(bcd)"
	r := MinRemoveToMakeValid(s)
	if r != "abcd(bcd)" {
		t.Errorf("must be a(b), is " + r)
	}
}
