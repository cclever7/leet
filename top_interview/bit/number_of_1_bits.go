package bit

// 位1的个数
// https://leetcode.cn/problems/number-of-1-bits/?envType=study-plan-v2&envId=top-interview-150
func hammingWeight(num uint32) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		if (1 << i & num) > 0 {
			cnt++
		}
	}
	return cnt
}
