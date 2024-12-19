package leetcode

import "fmt"

// still wrong

func LongestPalindrome() string {
	s := "cbbd"
	var stringLeft []rune
	var stringRight string

	var longestPalin int
	var longestString string

	for i := 0; i < len(s); i++ {
		holdString := string(s[i:])
		fmt.Println(stringRight)
		for h := 0; h < len(holdString); h++ {
			stringRight = holdString[0:h]
			stringLeft = []rune(holdString)
			if len(stringRight) > 1 {
				for x, j := 0, len(stringLeft)-1; x < j; x, j = x+1, j-1 {

					stringLeft[x], stringLeft[j] = stringLeft[j], stringLeft[x]
				}
			}
			if string(stringLeft) == stringRight && len(stringRight) > longestPalin {
				longestPalin = len(stringRight)
				longestString = stringRight
			}
		}
	}
	return longestString
}
