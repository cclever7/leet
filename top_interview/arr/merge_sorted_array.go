package arr

// 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {
	ret := make([]int, 0)
	nums1Idx, nums2Idx := 0, 0
	for {
		if nums1Idx >= m || nums2Idx >= n {
			break
		}
		if nums1[nums1Idx] < nums2[nums2Idx] {
			ret = append(ret, nums1[nums1Idx])
			nums1Idx++
		} else {
			ret = append(ret, nums2[nums2Idx])
			nums2Idx++
		}
	}
	if nums1Idx != m {
		for i := nums1Idx; i < m; i++ {
			ret = append(ret, nums1[i])
		}
	} else if nums2Idx != n {
		for i := nums2Idx; i < n; i++ {
			ret = append(ret, nums2[i])
		}
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = ret[i]
	}
}
