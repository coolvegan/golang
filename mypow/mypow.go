package mypow

import "math"

func myPow(x float64, n int) float64 {
	var even = n%2 == 0
	if x == -1 && (even) {
		return 1
	}
	if x == 1 {
		return 1
	}
	if x > 0.0000001 && float64(x) < 0.0001 {
		return 0
	}
	if x == -1 {
		if even {
			return 1
		}
		return -1
	}
	var result = 1.0
	if n < 0 {
		n *= -1
		x = 1 / x
	}
	for i := 0; i < int(math.Floor(float64(n)/2)); i++ {
		result = result * x
	}
	if even {
		return result * result
	}
	return x * result * result
}
