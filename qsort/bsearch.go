package main

func bsearch(data *[]int, left, right, element int) (int, bool) {
	if len(*data) == 0 {
		return -1, false
	}
	if (*data)[left] == element {
		return left, true
	}
	if (*data)[right] == element {
		return right, true
	}
	p3 := ((right - left) / 2) + left
	if left+1 == right {
		return -1, false
	}
	if p3 == left {
		p3++
	} else if p3 == right {
		p3--
	}
	if (*data)[p3] == element {
		return p3, true
	}
	if left+1 == right {
		return -1, false
	}
	if (*data)[p3] < element {
		left = p3
	} else {
		right = p3
	}
	return bsearch(data, left, right, element)
}

func bsearchIt(data *[]int, element int) (int, bool) {
	p1 := 0
	p2 := len(*data) - 1
	for p1+1 < p2 {
		p3 := (p2 - p1/2) + p1
		if (*data)[p1] == element {
			return p1, true
		}
		if (*data)[p2] == element {
			return p2, true
		}
		if p3 == p1 {
			p3++
		} else if p3 == p2 {
			p3--
		}
		if (*data)[p3] == element {
			return p3, true
		} else {
			if (*data)[p3] < element {
				p1 = p3
			} else {
				p2 = p3
			}
		}
	}
	return -1, false
}
func bsearchItIndexedOrg(data *[]int, left, right, element int) (int, bool) {
	p1 := left
	p2 := right
	for p1+1 < p2 {
		p3 := (p2 - p1/2) + p1
		if (*data)[p1] == element {
			return p1, true
		}
		if (*data)[p2] == element {
			return p2, true
		}
		if p3 == p1 {
			p3++
		} else if p3 == p2 {
			p3--
		}
		if (*data)[p3] == element {
			return p3, true
		} else {
			if (*data)[p3] < element {
				p1 = p3
			} else {
				p2 = p3
			}
		}
	}
	return -1, false
}
func bsearchItIndexed(data *[]int, left, right, element int) (int, bool) {
	p1 := left
	p2 := right
	for p1 <= p2 && p1 >= left && p2 <= right {
		p3 := p2 + p1/2
		if (*data)[p3] == element {
			return p3, true
		} else {
			if (*data)[p3] < element {
				p1 = p3 + 1
			} else {
				p2 = p3 - 1
			}
		}
	}
	return -1, false
}

func max(val1, val2 int) int {
	if val1 >= val2 {
		return val1
	}
	return val2
}
func min(val1, val2 int) int {
	if val1 <= val2 {
		return val1
	}
	return val2
}

func findStartAndEndIndexLinear(data *[]int, element int) (int, int, bool) {
	val, _ := bsearchItIndexed(data, 0, len(*data)-1, element)
	if val == -1 {
		return -1, -1, false
	}
	right := val
	for i := val + 1; i < len(*data); i++ {
		if (*data)[i] == element {
			right = max(right, i)
		}
	}
	left := val
	for i := val - 1; i >= 0; i-- {
		if (*data)[i] == element {
			left = min(left, i)
		}
	}
	return left, right, true
}

func findStartAndEndIndex(data *[]int, element int) (int, int, bool) {
	val, _ := bsearchItIndexed(data, 0, len(*data)-1, element)
	if val == -1 {
		return -1, -1, false
	}
	left := val
	right := val
	tmp := val
	for tmp >= 0 {
		tmp, _ = bsearchItIndexed(data, 0, tmp, element)
		if tmp != -1 {
			left = min(left, tmp)
		}
		tmp--
	}
	return left, right, true
}
