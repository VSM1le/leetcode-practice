package leetcode

func LongestCommonPrefix() string {
	strs := []string{"ab", "a"}
	if len(strs) == 1 {
		return strs[0]
	}
	iteration := true
	var count int
	var ans string
	for iteration {
		var c byte
		for _, v := range strs {
			if c == 0 {
				c = v[count]
			}
			if v[count] != 0 && c != v[count] {
				iteration = false
			}
		}

		if iteration {
			ans += string(c)
			count++
		}
	}
	return ans
}
