package leetcode

import "fmt"

func LengthOfLongestSubstring() int {
	varString := "abba"
	var lenght int
	var start int
	mapCheck := make(map[rune]int)
	for i, v := range varString {
		if lastIndex, ok := mapCheck[v]; ok && lastIndex >= start {
			start = mapCheck[v] + 1
			fmt.Println("start", start)
		}
		mapCheck[v] = i

		if ((i - start) + 1) > lenght {
			lenght = (i - start) + 1
			fmt.Println("start lenght", start)
		}
		fmt.Println("current length", lenght)
	}
	return lenght
}
