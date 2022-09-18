package main

import (
	"testing"
)

func TestCaseString1(t *testing.T) {
	mystr := "abc"
	mystr2 := "abc"
	var equal = compare(mystr, mystr2)
	if equal == false {
		t.Errorf("String must be abc, is true ")
	}

}

func TestCaseString2(t *testing.T) {
	mystr := "abc#"
	mystr2 := "abc"
	var equal = compare(mystr, mystr2)
	if equal == true {
		t.Errorf("Strings must be false  ")
	}

}

func TestCaseString3(t *testing.T) {
	mystr := "abc#"
	mystr2 := "ab"
	var equal = compare(mystr, mystr2)
	if equal == false {
		t.Errorf("Strings must be true, is false ")
	}

}

func TestCaseString4(t *testing.T) {
	mystr := "abc##"
	mystr2 := "ab#"
	var equal = compare(mystr, mystr2)
	if equal == false {
		t.Errorf("Strings must be true, is false ")
	}

}

func TestCaseString5(t *testing.T) {
	mystr := "abc##d"
	mystr2 := "ab#dd#"
	var equal = compare(mystr, mystr2)
	if equal == false {
		t.Errorf("Strings must be true, is false ")
	}

}

func TestCaseString6(t *testing.T) {
	mystr := "a#c"
	mystr2 := "b"
	var equal = compare(mystr, mystr2)
	if equal == true {
		t.Errorf("Strings must be false, is true ")
	}

}
func TestCaseString7(t *testing.T) {
	mystr := "xmp"
	mystr2 := "xrmu#p"
	var equal = compare(mystr, mystr2)
	if equal == false {
		t.Errorf("Strings must be true, is false ")
	}

}

func TestCaseString8(t *testing.T) {
	mystr := "bxj##tw"
	mystr2 := "bxj###tw"
	var equal = compare(mystr, mystr2)
	if equal == true {
		t.Errorf("Strings must be false, is true ")
	}

}
