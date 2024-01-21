package bit

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	ret := hammingWeight(00000000000000000000000000001011)
	fmt.Println(ret)
}

func Test_singleNumber(t *testing.T) {
	nums := []int{1, 2, 2, 3, 3}
	ret := singleNumber(nums)
	fmt.Println(ret)
}
