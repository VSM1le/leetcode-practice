package leetcode

import (
	"math"
)

func is32BitInteger(n int) bool {
	return n >= math.MinInt32 && n <= math.MaxInt32
}

func ReverseNumber() int {

	x := 1534236469
	converMinus := 1
	if x < 0 {
		x = x * -1
		converMinus = -1
	}
	var reverse int
	for x > 0 {
		reverse = (reverse * 10) + (x % 10)
		x = x / 10
	}
	reverse = reverse * converMinus
	if !is32BitInteger(reverse) {
		return 0
	}
	return reverse
}
