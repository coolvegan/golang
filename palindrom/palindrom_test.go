package palindrom

import "testing"

func TestCase1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	r := var1(s)

	if r == false {
		t.Errorf("Must be true")
	}

}

func TestCase2(t *testing.T) {
	s := ""
	r := var1(s)

	if r == true {
		t.Errorf("Empty must be false")
	}

}

func TestCase3(t *testing.T) {
	s := "a"
	r := var1(s)

	if r == false {
		t.Errorf("One character must be true")
	}

}

func TestCase4(t *testing.T) {
	s := "abc"
	r := var1(s)

	if r != false {
		t.Errorf("Not a palindrom must be false")
	}

}

func TestVarianteZwei1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	r := var2(s)

	if r == false {
		t.Errorf("Must be true")
	}

}

func TestVarianteZwei2(t *testing.T) {
	s := ""
	r := var2(s)

	if r == true {
		t.Errorf("Empty must be false")
	}

}

func TestVarianteZwei3(t *testing.T) {
	s := "a"
	r := var2(s)

	if r == false {
		t.Errorf("One character must be true")
	}

}

func TestVarianteZwei4(t *testing.T) {
	s := "abc"
	r := var2(s)

	if r != false {
		t.Errorf("Not a palindrom must be false")
	}

}

func TestReverse(t *testing.T) {
	s := "abc"
	r := reverse(s)

	if r != "cba" {
		t.Errorf("abc must be reversed cba")
	}

	s = "Marco"
	r = reverse(s)

	if r != "ocraM" {
		t.Errorf("abc must be reversed cba")
	}

	s = ""
	r = reverse(s)

	if r != "" {
		t.Errorf("abc must be reversed cba")
	}

	s = "a"
	r = reverse(s)

	if r != "a" {
		t.Errorf("abc must be reversed cba")
	}

}
func TestVariantDrei1(t *testing.T) {
	s := "A man, a plan, a canal: Panama"
	r := var2(s)

	if r == false {
		t.Errorf("Must be true")
	}

}

func TestVariantDrei2(t *testing.T) {
	s := ""
	r := var2(s)

	if r == true {
		t.Errorf("Empty must be false")
	}

}

func TestVariantDrei3(t *testing.T) {
	s := "a"
	r := var2(s)

	if r == false {
		t.Errorf("One character must be true")
	}

}

func TestVariantDrei4(t *testing.T) {
	s := "abc"
	r := var2(s)

	if r != false {
		t.Errorf("Not a palindrom must be false")
	}

}

func TestGetString(t *testing.T) {
	s := "a   b C381#!"
	r := getString(s)

	if r != "abc" {
		t.Errorf("Must be abc")
	}

}
