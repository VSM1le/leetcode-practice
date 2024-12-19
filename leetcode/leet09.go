package leetcode

func IsPalindrome() bool {
	par := 121
	copy := par
	var palinDrome int
	for copy > 0 {
		palinDrome = (palinDrome * 10) + (copy % 10)
		copy = copy / 10
	}

	return palinDrome == par
}
