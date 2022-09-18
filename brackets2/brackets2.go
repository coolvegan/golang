package main

import "strings"

type Stack []int

func (s *Stack) push(val int) {
	*s = append(*s, val)
}

func (s *Stack) pop() (int, bool) {
	if !(*s).isEmpty() {
		idx := len(*s) - 1
		elm := (*s)[idx]
		*s = (*s)[:idx]
		return elm, true
	}
	return 0, false
}

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func MinRemoveToMakeValid2(s string) string {
	var stck []int
	i := 0
	str := strings.Split(s, "")
	for i < len(str) {
		if len(stck) == 0 && string(str[i]) == ")" {
			str[i] = ""
			continue
		} else if len(stck) > 0 && string(str[i]) == ")" {
			stck = stck[:len(stck)-1]
		} else if string(str[i]) == "(" {
			stck = append(stck, i)
		}
		i++
	}
	for i := 0; i < len(stck); i++ {
		str[stck[i]] = ""
	}
	return strings.Join(str, "")
}

func MinRemoveToMakeValid(s string) string {
	var stack Stack
	i := 0
	for i < len(s) {
		if stack.isEmpty() && string(s[i]) == ")" {
			s = s[:i] + s[i+1:]
			continue
		} else if stack.isEmpty() != true && string(s[i]) == ")" {
			stack.pop()
		} else if string(s[i]) == "(" {
			stack.push(i)
		}
		i++
	}
	for stack.isEmpty() != true {
		i, _ := stack.pop()
		s = s[:i] + s[i+1:]
	}
	return s

}
