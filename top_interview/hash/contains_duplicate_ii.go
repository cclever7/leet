package hash

// 存在重复元素II
// https://leetcode.cn/problems/contains-duplicate-ii/description/?envType=study-plan-v2&envId=top-interview-1507
func containsNearbyDuplicate(nums []int, k int) bool {
	mp := make(map[int][]int)
	for idx, v := range nums {
		idxList, ok := mp[v]
		if !ok {
			mp[v] = make([]int, 0)
			mp[v] = append(mp[v], idx)
		} else {
			if idx-idxList[len(idxList)-1] <= k {
				return true
			} else {
				mp[v] = append(mp[v], idx)
			}
		}
	}
	return false
}
