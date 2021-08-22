// https://leetcode.com/problems/palindrome-number/
package main

import "fmt"

func isPalindrome(x int) bool {

	reverseNum := func(x int) int {
		revNum := 0
		for x > 0 {
			rem := x % 10
			revNum = revNum*10 + rem
			x = x / 10
		}
		return revNum
	}

	// for numbers less thn 1 return false
	if x < 0 {
		return false
	}

	// reverse the number and compare with original
	return x == reverseNum(x)

}

func main() {
	fmt.Println(0, isPalindrome(0))
	fmt.Println(5, isPalindrome(5))
	fmt.Println(10, isPalindrome(10))
	fmt.Println(121, isPalindrome(121))
	fmt.Println(-2323, isPalindrome(-2323))
	fmt.Println(123, isPalindrome(123))
}
