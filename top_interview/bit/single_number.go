package bit

// 只出现一次的数字
// https://leetcode.cn/problems/single-number/description/?envType=study-plan-v2&envId=top-interview-150
func singleNumber(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	ret := nums[0]
	for i := 1; i < length; i++ {
		ret = ret ^ nums[i]
	}
	return ret
}
