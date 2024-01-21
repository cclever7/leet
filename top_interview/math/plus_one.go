package math

// 加1
// https://leetcode.cn/problems/plus-one/description/?envType=study-plan-v2&envId=top-interview-150
func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{1}
	}

	carry := 0
	for i := length - 1; i >= 0; i-- {
		curr := 0
		if i == length-1 {
			curr = digits[i] + 1 + carry
		} else {
			curr = digits[i] + carry
		}
		carry = 0 // reset carry
		if curr < 10 {
			digits[i] = curr
		} else {
			digits[i] = curr % 10
			carry = 1
		}
		if carry != 1 {
			break
		}
	}
	if carry == 1 {
		ret := append([]int{1}, digits...)
		return ret
	}
	return digits
}
