package link

// 环形链表
// https://leetcode.cn/problems/linked-list-cycle/?envType=study-plan-v2&envId=top-interview-150
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	mp := make(map[*ListNode]struct{})
	for {
		if head.Next == nil {
			break
		}
		_, ok := mp[head]
		if !ok {
			mp[head] = struct{}{}
		} else {
			return true
		}
		head = head.Next
	}
	return false
}
