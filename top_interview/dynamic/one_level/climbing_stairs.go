package one_level

// 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/?envType=study-plan-v2&envId=top-interview-150
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	arr := make([]int, n)
	arr[0] = 1
	arr[1] = 2
	for i := 2; i < n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	return arr[n-1]
}
