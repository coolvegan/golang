package main

func compare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1
	for p1 >= 0 || p2 >= 0 {
		if (p1 < 0) || (p2 < 0) {
			return false
		}
		if s[p1] == '#' || t[p2] == '#' {

			if p1 >= 0 && s[p1] == '#' {
				backspace := 2
				for backspace > 0 {
					p1--
					backspace--
					if p1 >= 0 && s[p1] == '#' {
						backspace += 2
					}
				}
			}
			if p2 >= 0 && t[p2] == '#' {
				backspace := 2
				for backspace > 0 {
					p2--
					backspace--
					if p2 >= 0 && t[p2] == '#' {
						backspace += 2
					}
				}
			}
		} else {
			if (p1 < 0) || (p2 < 0) {
				return false
			}
			if s[p1] != t[p2] {
				return false
			} else {
				p2--
				p1--
			}
		}
	}
	return true
}
