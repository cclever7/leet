package top_interview

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	ret := removeElement(nums, 3)
	fmt.Println(nums, ret)
}

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 2, 2, 3}
	ret := removeDuplicates(nums)
	fmt.Println(nums, ret)
}

func Test_majorityElement(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	ret := majorityElement(nums)
	fmt.Println(nums, ret)
}
