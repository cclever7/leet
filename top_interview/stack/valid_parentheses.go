package stack

// 有效的括号
// https://leetcode.cn/problems/valid-parentheses/?envType=study-plan-v2&envId=top-interview-150
func isValid(s string) bool {
	stack := make([]rune, 0)
	var sRune = []rune(s)
	for _, v := range sRune {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack)-1 < 0 {
				return false
			}
			last := stack[len(stack)-1]
			switch v {
			case ')':
				if last != '(' {
					return false
				}
			case ']':
				if last != '[' {
					return false
				}
			case '}':
				if last != '{' {
					return false
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
