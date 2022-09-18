package palindrom

import (
	"regexp"
	"strings"
)

func getString(s string) string {
	m := regexp.MustCompile("[^a-zA-Z]")
	res := m.ReplaceAllString(strings.ToLower(s), "")
	return res
}

func validSubPalindrom(s string, p1, p2 int) bool {
	if p1 > p2 {
		return false
	}
	for p1 < p2 {
		if s[p1] != s[p2] {
			return false
		}

		p1++
		p2--
	}
	return true
}

func varextented(s string) bool {
	s = getString(s)
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	p1 := 0
	p2 := len(s) - 1

	for p1 < p2 {
		if s[p1] != s[p2] {
			return validSubPalindrom(s, p1+1, p2) || validSubPalindrom(s, p1, p2-1)

		}
		p1++
		p2--
	}
	return true
}

func var1(s string) bool {
	s = getString(s)
	if len(s) == 0 {
		return false
	}

	if len(s) == 1 {
		return true
	}

	p1 := 0
	p2 := len(s) - 1

	for p1 < p2 {
		if s[p1] != s[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

func var2(s string) bool {
	s = getString(s)
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}

	p1 := 0
	p2 := len(s) - 1

	length := len(s)
	isEvenString := length % 2
	if isEvenString == 0 {
		p1 = length/2 - 1
		p2 = length / 2
	} else {
		p1 = length / 2
		p2 = length / 2
	}

	for p1 >= 0 && p2 < length {
		if s[p1] != s[p2] {
			return false
		}
		p1--
		p2++
	}

	return true
}

func reverse(s string) string {
	resultStr := ""
	for i := len(s) - 1; i >= 0; i-- {
		resultStr += string((s[i]))
	}
	return resultStr
}

func var3(s string) bool {
	s = getString(s)
	rs := reverse(s)
	if len(rs) != len(s) {
		return false
	}
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	p1 := 0
	p2 := 0

	for p1 < len(s) && p2 < len(rs) {
		if s[p1] != rs[p2] {
			return false
		}
		p1++
		p2++
	}

	return true
}
