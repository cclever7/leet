package arr

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	var (
		i            = 0
		j            = i + 1
		duplicateCnt = 0
	)

	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
			duplicateCnt++
		} else {
			nums[i+1] = nums[j]
			i++
			j++
		}
	}
	return len(nums) - duplicateCnt
}
