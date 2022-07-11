package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var input1, input2, input3, input4, input5, input6 ListNode
	// var input1, input2, input3, input4 ListNode
	input1 = ListNode{Val: 1, Next: &input2}
	input2 = ListNode{Val: 2, Next: &input3}
	input3 = ListNode{Val: 3, Next: &input4}
	input4 = ListNode{Val: 4, Next: &input5}
	input5 = ListNode{Val: 5, Next: &input6}
	input6 = ListNode{Val: 6}
	// input4 = ListNode{Val: 4}

	result := middleNode(&input1)
	fmt.Println(result.Val)

	resultSlice := middleNodeSlice(&input1)
	fmt.Println(resultSlice.Val)

}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var fast, slow = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func middleNodeSlice(head *ListNode) *ListNode {
	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	fmt.Println(len(nodes))
	return nodes[len(nodes)/2]
}
