package math

import "strconv"

// 回文数
// https://leetcode.cn/problems/palindrome-number/description/?envType=study-plan-v2&envId=top-interview-150
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i != j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
		if i+1 == j {
			break
		}
	}
	return true
}
