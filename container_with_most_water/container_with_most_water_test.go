package solve

import (
	"strconv"
	"testing"
)

func Test_case1(t *testing.T) {
	var field = []int{2, 1, 0, 3}
	var r = solve(&field)
	if r != 3 {
		t.Errorf("Result must be 3 and is:" + strconv.Itoa(r))
	}
}

func Test_case2(t *testing.T) {
	var field = []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	var r = solve(&field)
	if r != 6 {
		t.Errorf("Result must be 6 and is:" + strconv.Itoa(r))
	}
}
func Test_case3(t *testing.T) {
	var field = []int{0, 2, 0, 1, 0, 3}
	var r = solve(&field)
	if r != 5 {
		t.Errorf("Result must be 5 and is:" + strconv.Itoa(r))
	}
}
