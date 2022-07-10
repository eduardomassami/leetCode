package main

import (
	"fmt"
	"reflect"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var input1, input2, input3, input4 ListNode
	input1 = ListNode{Val: 1, Next: &input2}
	input2 = ListNode{Val: 2, Next: &input3}
	input3 = ListNode{Val: 2, Next: &input4}
	input4 = ListNode{Val: 1}

	fmt.Println(isPalindrome(&input1))
}

func reverseLinkedList(head *ListNode, reverse *[]int) {
	if head == nil {
		return
	}

	reverseLinkedList(head.Next, reverse)
	*reverse = append(*reverse, head.Val)
	return
}

func isPalindrome(head *ListNode) bool {
	var current []int
	var reverse []int

	for curr := head; curr != nil; curr = curr.Next {
		current = append(current, curr.Val)
	}

	reverseLinkedList(head, &reverse)

	return reflect.DeepEqual(current, reverse)
}
