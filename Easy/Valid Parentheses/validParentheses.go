package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}
	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) < 1 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(top)
			fmt.Println(stack)
			if char == ')' && top != '(' || char == '}' && top != '{' || char == ']' && top != '[' {
				return false
			}
		}

	}
	return len(stack) == 0
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
