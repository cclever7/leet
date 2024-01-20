package hash

import "strings"

// 单词规律
// https://leetcode.cn/problems/word-pattern/?envType=study-plan-v2&envId=top-interview-150
func wordPattern(pattern string, s string) bool {
	var (
		sArr  = strings.Split(s, " ")
		mp1   = make(map[int32]string)
		mp2   = make(map[string]int32)
		match = true
	)
	if len(pattern) != len(sArr) {
		return false
	}
	for idx, v := range pattern {
		word, ok := mp1[v]
		if !ok {
			mp1[v] = sArr[idx]
		} else {
			if word != sArr[idx] {
				match = false
				break
			}
		}

		p, ok := mp2[sArr[idx]]
		if !ok {
			mp2[sArr[idx]] = v
		} else {
			if p != v {
				match = false
				break
			}
		}
	}
	return match
}
