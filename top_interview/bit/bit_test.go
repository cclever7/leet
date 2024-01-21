package bit

import (
	"fmt"
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	ret := hammingWeight(00000000000000000000000000001011)
	fmt.Println(ret)
}
