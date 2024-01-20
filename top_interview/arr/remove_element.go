package arr

// 移除元素
// https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
