package main

import "fmt"

// Convert int to str
/*
func isPalindrome(x int) bool {
	newString := strconv.Itoa(x)
	array := []string{}
	for _, v := range newString {
		array = append(array, string(v))
	}
	return array[len(array)-1] == array[0]
}

func main() {
	x := 11
	fmt.Println(isPalindrome(x))
}
*/

// Without convert int to str

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	/*
		reversed = 0
		temp = 121

		digit = 121 % 10 = 1
		reversed = 0 * 10 + 1 = 1
		temp = 121 / 10 = 12

		digit = 12 % 10 = 2
		reversed = 1 * 10 + 2 = 12
		temp = 12 / 10 = 1

		digit = 1 % 10 = 1
		reversed = 12 * 10 + 1 = 121
		temp = 1 / 10 = 0 -> temp == 0 -> break looping
	*/
	var reversed int = 0
	var temp int = x

	for temp != 0 {
		var digit int = temp % 10
		reversed = reversed*10 + digit // until temp = 0 then reversed is complete
		temp /= 10
	}
	return reversed == x
}

func main() {
	x := 121
	fmt.Println(isPalindrome(x))
}
