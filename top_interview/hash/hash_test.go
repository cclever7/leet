package hash

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	s1, s2 := "aa", "aab "
	fmt.Println(canConstruct(s1, s2))
}

func Test_isIsomorphic(t *testing.T) {
	s1, s2 := "paper", "title"
	fmt.Println(isIsomorphic(s1, s2))
}

func Test_wordPatternwordPattern(t *testing.T) {
	s1, s2 := "abba", "dog dog dog dog"
	fmt.Println(wordPattern(s1, s2))
}

func Test_twoSum(t *testing.T) {
	nums := []int{3, 2, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}

func Test_containsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 0, 1, 1}
	target := 1
	fmt.Println(containsNearbyDuplicate(nums, target))
}

func Test_isAnagram(t *testing.T) {
	s, v := "rat", "car"
	fmt.Println(isAnagram(s, v))
}
