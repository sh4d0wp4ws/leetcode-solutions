package main

import (
	"fmt"
)

func romanToInt(s string) int {
	if len(s) < 1 || len(s) > 15 {
		return 0
	}
	result := 0
	romanDefine := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	/*
		Input: "IV"
		First Loop:
		i = 0;
		value = romanDefine["I"] = 1
		if 1 < 2 && 1 < 5 (romanDefine[s[0+1]] = 5) -- true && true => result = 0 - 1 = -1

		Second Loop:
		i = 1;
		value = romanDefine("V") = 5
		if 2 < 2 && 5 < 10 (romanDefine[s[1+1]] = 10) -- false && false => result = -1 + 5 = 4

		Third Loop:
		i = 2; because len(s) = 2 => not meet condition => break the loop => return the result
	*/
	for i := 0; i < len(s); i++ {
		value := romanDefine[s[i]]
		if i+1 < len(s) && value < romanDefine[s[i+1]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}

func main() {
	s := "IV"
	fmt.Println(romanToInt(s))
}
