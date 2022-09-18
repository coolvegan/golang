package main

type Stack []string

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) *Stack {
	*s = append(*s, str)
	return s
}

func (s *Stack) Pop() (string, bool) {
	if s.isEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func brackets(str *string) bool {
	var stack Stack
	m := make(map[string]string)
	m["{"] = "}"
	m["("] = ")"
	m["["] = "]"
	for i := 0; i < len(*str); i++ {
		if len(m[string((*str)[i])]) != 0 {
			stack.Push(m[string((*str)[i])])
		} else {
			chr, _ := stack.Pop()
			if chr != string((*str)[i]) {
				return false
			}
		}
	}
	return stack.isEmpty()
}
