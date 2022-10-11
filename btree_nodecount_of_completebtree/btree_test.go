package countProblem

import "testing"

func TestPow(t *testing.T) {
	r := pow(2, 2)
	r2 := pow(2, 3)
	r3 := pow(2, 10)

	if r != 4 {
		t.Errorf("Must be four")
	}

	if r2 != 8 {
		t.Errorf("Must be 8")
	}
	if r3 != 1024 {
		t.Errorf("Must be 1024")
	}
}

func TestNodeExists(t *testing.T) {
	var a, b, c, d, e Node
	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e

	r := countNodes(&a)
	if r != 5 {
		t.Errorf("Must be 5, is %d", r)
	}

}
func TestNodeExists2(t *testing.T) {
	var a, b, c, d, e, f, g Node
	a.left = &b
	a.right = &c
	b.left = &d
	b.right = &e
	c.left = &f
	c.right = &g

	r := countNodes(&a)
	if r != 7 {
		t.Errorf("Must be 7, is %d", r)
	}

}
func TestNodeExists3(t *testing.T) {
	var a, b, c, d Node
	a.left = &b
	a.right = &c
	b.left = &d

	r := countNodes(&a)
	if r != 4 {
		t.Errorf("Must be 4, is %d", r)
	}
}

func TestIsValidBstCase1(t *testing.T) {
	var a, b, c, d, e Node
	a.left = &b
	a.right = &c
  b.left = &d
  b.right = &e

	a.val = 5
	b.val = 2
	c.val = 7
	d.val = 1
	e.val = 4

	r := isValidBST(&a)
	if r != true {
		t.Errorf("Must be true")
	}

}
func TestIsValidBstCase2(t *testing.T) {
	var a, b, c, d, e Node
	a.left = &b
	a.right = &c
  b.left = &d
  b.right = &e
	a.val = 5
	b.val = 2
	c.val = 7
	d.val = 6
	e.val = 4
	r := isValidBST(&a)
	if r == true {
		t.Errorf("Must be false")
	}
}
func TestIsValidBstCase3(t *testing.T) {
	var a, b, c, d, e Node
	a.left = &b
	a.right = &c
  b.left = &d
  b.right = &e
	a.val = 5
	b.val = 2
	c.val = 7
	d.val = 6
	e.val = 4
	r := isValidBST(&a)
	if r == true {
		t.Errorf("Must be false")
	}
}
func TestIsValidBstCase4(t *testing.T) {
	var a, b, c, d, e, f Node
	a.left = &b
	a.right = &c
  b.left = &d
  b.right = &e
  c.right = &f
	a.val = 5
	b.val = 2
	c.val = 7
	d.val = 6
	e.val = 4
  f.val = 1
	r := isValidBST(&a)
	if r == true {
		t.Errorf("Must be true")
	}
}
