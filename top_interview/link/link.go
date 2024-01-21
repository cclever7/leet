package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertArrayToLink(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	move := head
	for i := 1; i < len(nums); i++ {
		move.Next = &ListNode{Val: nums[i], Next: nil}
		move = move.Next
	}
	return head
}

func printfLink(head *ListNode) {
	if head == nil {
		fmt.Println("link is nil")
		return
	}
	for head != nil {
		fmt.Printf("%d\t", head.Val)
		head = head.Next
	}
}
