package main

func max(v1, v2 int) int {
	if v1 >= v2 {
		return v1
	}
	return v2
}

func min(v1, v2 int) int {
	if v2 <= v1 {
		return v2
	}
	return v1
}

func maxLeft(pos int, arr *[]int, valMap map[int]int) int {
	r := 0
	if valMap[pos] != 0 {
		return valMap[pos]
	}
	for i := pos; i >= 0; i-- {
		r = max(r, (*arr)[i])
	}
	valMap[pos] = r
	return r
}

func maxRight(pos int, arr *[]int, valMap map[int]int) int {
	r := 0
	if valMap[pos+1] != 0 {
		return valMap[pos+1]
	}
	for i := pos + 1; i < len(*arr); i++ {
		r = max(r, (*arr)[i])
	}
	valMap[pos+1] = r
	return r
}

func solve(arr *[]int) int {
	result := 0
	var maxLmap = make(map[int]int)
	var maxRmap = make(map[int]int)
	for i := 0; i < len(*arr); i++ {
		r := min(maxLeft(i, arr, maxLmap), maxRight(i, arr, maxRmap)) - (*arr)[i]
		if r > 0 {
			result += r
		}
	}
	return result
}
















func solve2(arr []int) int {
	result := 0
	p1 := 0
	p2 := len(arr) - 1
	maxl := 0
	maxr := 0
	for p1 < p2 {
		maxl = max(maxl, arr[p1])
		maxr = max(maxr, arr[p2])
		val := 0
		if maxl <= maxr {
			val = maxl - arr[p1]
			p1++
		} else {
			val = maxr - arr[p2]
			p2--
		}
		if val > 0 {
			result += val
		}
	}
	return result
}















func main() {
	var field = []int{0, 2, 0, 1, 0, 3}
	var r = solve(&field)
	max(r, r)
}
