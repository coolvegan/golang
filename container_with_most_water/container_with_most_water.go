package solve

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}
func solve(arr *[]int) int {
	result := 0
	p1 := 0
	p2 := len(*arr) - 1

	maxl := 0
	maxr := 0
	for p1 < p2 {
		maxl = max(maxl, (*arr)[p1])
		maxr = max(maxr, (*arr)[p2])
		val := 0
		if maxl <= maxr {
			val = maxl - (*arr)[p1]
			p1++
		} else {
			val = maxr - (*arr)[p2]
			p2--
		}

		if val > 0 {
			result += val
		}
	}
	return result
}
