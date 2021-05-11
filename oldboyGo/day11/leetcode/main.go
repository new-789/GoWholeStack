package main

import (
	"fmt"
)

// LeetCode
type ListNode struct {
	Val int
	Next *ListNode
}

func reverList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(head)
	ret := reverList(head)
	fmt.Println(ret)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
