package hash

// 赎金信
// https://leetcode.cn/problems/ransom-note/description/?envType=study-plan-v2&envId=top-interview-150
func canConstruct(ransomNote string, magazine string) bool {
	mp := make(map[int32]int)
	for _, v := range magazine {
		_, ok := mp[v]
		if !ok {
			mp[v] = 1
		} else {
			mp[v]++
		}
	}

	ret := true
	for _, v := range ransomNote {
		_, ok := mp[v]
		if !ok {
			ret = false
			break
		} else {
			mp[v]--
			if mp[v] < 0 {
				ret = false
				break
			}
		}
	}
	return ret
}
