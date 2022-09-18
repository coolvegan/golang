package sequence

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	var s = "abcb"
	p := predecessor(s, 3, 0)

	if p != 1 {
		t.Errorf("Must be one")
		fmt.Printf("p is %d", p)
	}
}

func TestMain2(t *testing.T) {
	var s = "abcb"
	p := predecessor(s, 2, 0)

	if p != -1 {
		t.Errorf("Must be zero")
	}
}

func TestMain3(t *testing.T) {
	var s = "aba"
	p := predecessor(s, 2, 0)

	if p != 0 {
		t.Errorf("Must be 1")
		fmt.Printf("p is %d", p)
	}
}

func TestMain4(t *testing.T) {
	var s = "efadef"
	p := predecessor(s, 5, 2)

	if p != -1 {
		t.Errorf("Must be 4")
		fmt.Printf("p is %d", p)
	}
}

func TestSubstringCount(t *testing.T) {
	var s = "abcabde"
	p := lengthOfLongestSubstring(s)

	if p != 5 {
		t.Errorf("Must be 5")
		fmt.Printf("p is %d", p)
	}
}

func TestSubstringCount2(t *testing.T) {
	var s = "aaaaa"
	p := lengthOfLongestSubstring(s)

	if p != 1 {
		t.Errorf("Must be 1")
		fmt.Printf("p is %d", p)
	}
}

func TestSubstringCount3(t *testing.T) {
	var s = ""
	p := lengthOfLongestSubstring(s)

	if p != 0 {
		t.Errorf("Must be 0")
		fmt.Printf("p is %d", p)
	}
}
func TestSubstringCount4(t *testing.T) {

	var s = "abcdecghij"
	p := lengthOfLongestSubstring(s)

	if p != 7 {
		t.Errorf("Must be 7")
		fmt.Printf("p is %d", p)
	}
}
func TestSubstringCount5(t *testing.T) {

	var s = "pwwkew"
	p := lengthOfLongestSubstring(s)

	if p != 3 {
		t.Errorf("Must be 3")
		fmt.Printf("p is %d", p)
	}
}

func TestSubstringCount6(t *testing.T) {

	var s = "au"
	p := lengthOfLongestSubstring(s)

	if p != 2 {
		t.Errorf("Must be 2")
		fmt.Printf("p is %d", p)
	}
}

func TestSubstringCount7(t *testing.T) {

	var s = "dvdf"
	p := lengthOfLongestSubstring(s)

	if p != 3 {
		t.Errorf("Must be 3")
		fmt.Printf("p is %d", p)
	}
}

func TestHashMap1(t *testing.T) {
	var cache = make(map[rune]int)
	cache['a'] = 5
	if cache['a'] != 5 {
		t.Errorf("Size of letter a must be 5")
		fmt.Printf("is %d", cache['a'])
	}
}
