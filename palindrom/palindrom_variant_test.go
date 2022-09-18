package palindrom

import "testing"

func TestRaceACar(t *testing.T) {
	s := "raceacar"
	r := varextented(s)

	if r == false {
		t.Errorf("Must be true")
	}

}

func TestAbccdba(t *testing.T) {
	s := "abccdba"
	r := varextented(s)

	if r == false {
		t.Errorf("Must be true")
	}

}

func TestAbcdefdba(t *testing.T) {
	s := "abcdefdba"
	r := varextented(s)

	if r == true {
		t.Errorf("Must be false")
	}

}
