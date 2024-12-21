package leetcode

import "fmt"

func IsValid() bool {
	s := "()"
	mapChar := map[rune]rune{']': '[', '}': '{', ')': '('}
	var stack []rune
	for _, v := range s {
		if string(v) == "[" || string(v) == "{" || string(v) == "(" {
			stack = append(stack, v)
			continue
		}

		if len(stack) == 0 {
			return false
		} else if mapChar[v] == stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		} else {
			return false
		}

	}
	fmt.Println(stack, len(stack))
	return len(stack) == 0
}
