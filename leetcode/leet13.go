package leetcode

func RomanToInt() int {
	s := "LVIII"
	mapRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var value int
	for i := 0; i < len(s)-1; i++ {
		if mapRoman[s[i]] < mapRoman[s[i+1]] {
			value -= mapRoman[s[i]]
		} else {
			value += mapRoman[s[i]]
		}
	}
	return value + mapRoman[s[len(s)-1]]
}
