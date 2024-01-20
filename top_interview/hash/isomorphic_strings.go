package hash

// 同构字符串
// https://leetcode.cn/problems/isomorphic-strings/?envType=study-plan-v2&envId=top-interview-150
func isIsomorphic(s string, t string) bool {
	return realIs(s, t) && realIs(t, s)
}

func realIs(s string, t string) bool {
	isomorphic := true
	mp := make(map[uint8]uint8)
	for i := 0; i < len(s); i++ {
		tEle, ok := mp[s[i]]
		if !ok {
			mp[s[i]] = t[i]
		} else {
			if tEle != t[i] {
				isomorphic = false
			}
		}
	}
	return isomorphic
}
