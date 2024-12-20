package leetcode

func MaxArea() int {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	res := 0
	ans := 0
	l, r := 0, len(height)-1
	for l != r {
		if height[l] < height[r] {
			res = (r - l) * height[l]
			l += 1
		} else {
			res = (r - l) * height[r]
			r -= 1
		}
		if res > ans {
			ans = res
		}
	}
	return ans
}
