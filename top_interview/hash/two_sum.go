package hash

// 两数之和
// https://leetcode.cn/problems/two-sum/description/?envType=study-plan-v2&envId=top-interview-150
func twoSum(nums []int, target int) []int {
	prevNums := make(map[int]int)
	for i, num := range nums {
		rest := target - num
		restIdx, ok := prevNums[rest]
		if ok {
			return []int{restIdx, i}
		} else {
			prevNums[num] = i
		}
	}
	return []int{}
}
