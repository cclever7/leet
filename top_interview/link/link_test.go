package link

import (
	"fmt"
	"testing"
)

func Test_convertArrayToLink(t *testing.T) {
	nums := []int{1}
	head := convertArrayToLink(nums)
	printfLink(head)
}

func Test_hasCycle(t *testing.T) {
	link := convertArrayToLink([]int{3, 2, 0, -4})
	fmt.Println(hasCycle(link))
}
