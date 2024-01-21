package hash

// 有效字母异位词
// https://leetcode.cn/problems/valid-anagram/description/?envType=study-plan-v2&envId=top-interview-150
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[int32]int)
	for _, v := range s {
		mp[v]++
	}
	for _, v := range t {
		_, ok := mp[v]
		if !ok {
			return false
		} else {
			mp[v]--
			if mp[v] < 0 {
				return false
			}
		}
	}
	return true
}
