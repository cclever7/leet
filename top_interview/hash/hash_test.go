package hash

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	s1, s2 := "aa", "aab "
	fmt.Println(canConstruct(s1, s2))
}
