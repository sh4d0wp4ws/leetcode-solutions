package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, value := range strs {
		for len(strs) > 0 {
			if strings.HasPrefix(value, prefix) {
				break
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	strs := []string{"flower", "fkow"}
	fmt.Println(longestCommonPrefix(strs))
}
