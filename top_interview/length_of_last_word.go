package top_interview

// 最后一个单词的长度
// https://leetcode.cn/problems/length-of-last-word/?envType=study-plan-v2&envId=top-interview-150
func lengthOfLastWord(s string) int {
	idx := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			idx = i
			break
		}
	}
	length := 0
	for i := idx; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		} else {
			break
		}
	}
	return length
}
