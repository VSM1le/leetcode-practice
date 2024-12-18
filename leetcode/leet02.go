package leetcode

import "fmt"

// not pass all testcase
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseNumber(number int) int {
	var reverseNumber int
	for number > 0 {
		reverseNumber = (reverseNumber * 10) + (number % 10)
		number = number / 10
	}
	return reverseNumber
}

func ConverLinkListToNumber(list *ListNode) *int {
	var valueList int
	for list != nil {
		valueList = (valueList * 10) + list.Val
		list = list.Next // Traverse the list
	}
	return &valueList
}

func Insert(l1 *ListNode, l2 *ListNode) *ListNode {
	valueList1 := ConverLinkListToNumber(l1)
	valueList2 := ConverLinkListToNumber(l2)
	reverse1 := reverseNumber(*valueList1)
	reverse2 := reverseNumber(*valueList2)
	sumReverse := reverse1 + reverse2
	fmt.Println(sumReverse)
	if sumReverse < 10 {
		fmt.Println("hello")
		return &ListNode{Val: sumReverse}
	}

	var head *ListNode
	var tail *ListNode
	for sumReverse > 0 {
		newNode := &ListNode{Val: sumReverse % 10}
		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.Next = newNode
			tail = newNode
		}
		sumReverse = sumReverse / 10
	}
	return head
}
